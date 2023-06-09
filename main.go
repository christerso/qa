package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var db *sql.DB

func init() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	host := os.Getenv("STEPQA_DB_HOST")
	if host == "" {
		host = os.Getenv("STEPQA_DB_HOST")
		if host == "" {
			log.Fatal("STEPQA_DB_HOST environment variable not set")
		}
	}

	username := os.Getenv("STEPQA_DB_USERNAME")
	if username == "" {
		username = os.Getenv("STEPQA_DB_USERNAME")
		if username == "" {
			log.Fatal("STEPQA_DB_USERNAME environment variable not set")
		}
	}

	password := os.Getenv("STEPQA_DB_PASSWORD")
	if password == "" {
		password = os.Getenv("STEPQA_DB_PASSWORD")
		if password == "" {
			log.Fatal("STEPQA_DB_PASSWORD environment variable not set")
		}
	}

	database := os.Getenv("STEPQA_DB_NAME")
	if database == "" {
		database = os.Getenv("STEPQA_DB_NAME")
		if database == "" {
			log.Fatal("STEPQA_DB_NAME environment variable not set")
		}
	}
	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, host, database)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

type QA struct {
	ID        int
	Topic     string
	Subtopic  string
	Question  string
	Answer    string
	Steps     []string
	DoneSteps []int
}

func printLogo() {
	logo := `### StepQA ###                   
`
	fmt.Println(logo)
}

func listAll() ([]QA, error) {
	rows, err := db.Query("SELECT id, topic, subtopic, question, answer, steps FROM questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allQA []QA
	for rows.Next() {
		var q QA
		var steps string
		err := rows.Scan(&q.ID, &q.Topic, &q.Subtopic, &q.Question, &q.Answer, &steps)
		if err != nil {
			return nil, err
		}
		json.Unmarshal([]byte(steps), &q.Steps)
		allQA = append(allQA, q)
	}
	return allQA, nil
}

func inputQAInfo(existingQA QA) QA {
	topicPrompt := promptui.Prompt{
		Label:   "Topic",
		Default: existingQA.Topic,
	}
	topic, _ := topicPrompt.Run()

	subtopicPrompt := promptui.Prompt{
		Label:   "Subtopic",
		Default: existingQA.Subtopic,
	}
	subtopic, _ := subtopicPrompt.Run()

	questionPrompt := promptui.Prompt{
		Label:   "Question",
		Default: existingQA.Question,
	}
	question, _ := questionPrompt.Run()

	answerPrompt := promptui.Prompt{
		Label:   "Answer",
		Default: existingQA.Answer,
	}
	answer, _ := answerPrompt.Run()

	steps := existingQA.Steps
	for {
		stepPrompt := promptui.Prompt{
			Label: "Add a step (leave blank when done)",
		}
		step, _ := stepPrompt.Run()
		if step == "" {
			break
		}
		steps = append(steps, step)
	}

	return QA{
		ID:       existingQA.ID,
		Topic:    topic,
		Subtopic: subtopic,
		Question: question,
		Answer:   answer,
		Steps:    steps,
	}
}

func displayQAList(qaList []QA, checkmarks map[int][]int) {
	for _, qa := range qaList {
		color.New(color.FgYellow).Printf("\nID: %d\n", qa.ID)
		fmt.Printf("Topic: %s\n", qa.Topic)
		fmt.Printf("Subtopic: %s\n", qa.Subtopic)
		fmt.Printf("Question: %s\n", qa.Question)
		fmt.Printf("Answer: %s\n", qa.Answer)

		if len(qa.Steps) > 0 {
			fmt.Printf("Steps:\n")
			for i, step := range qa.Steps {
				done := false
				for _, doneStep := range checkmarks[qa.ID] {
					if doneStep == i+1 { // Adjust the step index
						done = true
						break
					}
				}
				if done {
					color.New(color.FgGreen).Printf("\t%d. [✓] %s\n", i+1, step)
				} else {
					color.New(color.FgRed).Printf("\t%d. [ ] %s\n", i+1, step)
				}
			}
		}
	}
}

func searchQA(searchTerm string) ([]QA, error) {
	rows, err := db.Query("SELECT id, topic, subtopic, question, answer, steps FROM questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allQA []QA
	for rows.Next() {
		var q QA
		var steps string
		err := rows.Scan(&q.ID, &q.Topic, &q.Subtopic, &q.Question, &q.Answer, &steps)
		if err != nil {
			return nil, err
		}
		json.Unmarshal([]byte(steps), &q.Steps)
		allQA = append(allQA, q)
	}

	var results []QA
	lowerSearchTerm := strings.ToLower(searchTerm)
	for _, q := range allQA {
		if strings.Contains(strings.ToLower(q.Topic), lowerSearchTerm) || strings.Contains(strings.ToLower(q.Subtopic), lowerSearchTerm) || strings.Contains(strings.ToLower(q.Question), lowerSearchTerm) {
			results = append(results, q)
		}
	}
	return results, nil
}

func toggleCompletion(checkmarks map[int][]int) {
	idPrompt := promptui.Prompt{
		Label: "Enter the ID of the Q&A",
	}
	idStr, err := idPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}
	id, _ := strconv.Atoi(idStr)

	qa, err := getQAByID(id)
	if err != nil {
		fmt.Printf("Error fetching Q&A by ID: %v\n", err)
		return
	}

	stepItems := make([]string, len(qa.Steps))

	updateStepItems := func() {
		for i, step := range qa.Steps {
			checked := ""
			for _, doneStep := range checkmarks[id] {
				if doneStep == i+1 {
					checked = "✓ "
					break
				}
			}
			stepWithoutNumber := strings.Join(strings.Split(step, " ")[1:], " ")
			stepItems[i] = fmt.Sprintf("%d. %s%s", i+1, checked, stepWithoutNumber)
		}
	}

	updateStepItems()

	stepPrompt := promptui.Select{
		Label:        "Select a step to toggle (press Ctrl-C to finish)",
		Items:        stepItems,
		HideSelected: true,
	}

	for {
		_, stepStr, err := stepPrompt.Run()
		if err != nil {
			if err == promptui.ErrAbort {
				break
			} else {
				fmt.Printf("Prompt failed: %v\n", err)
				return
			}
		}

		stepIndexStr := strings.Split(stepStr, ".")[0]
		stepIndex, err := strconv.Atoi(stepIndexStr)
		if err != nil {
			fmt.Printf("Error converting step index to integer: %v\n", err)
			return
		}

		if checkmarks[id] == nil {
			checkmarks[id] = []int{}
		}

		found := false
		for i, doneStep := range checkmarks[id] {
			if doneStep == stepIndex {
				checkmarks[id] = append(checkmarks[id][:i], checkmarks[id][i+1:]...)
				found = true
				break
			}
		}

		if !found {
			checkmarks[id] = append(checkmarks[id], stepIndex)
		}

		updateStepItems()
	}
}

func getQAByID(id int) (QA, error) {
	allQA, err := listAll()
	if err != nil {
		return QA{}, err
	}

	for _, qa := range allQA {
		if qa.ID == id {
			return qa, nil
		}
	}

	return QA{}, fmt.Errorf("Q&A not found for ID: %d", id)
}

func addQA(qa QA) error {
	steps, _ := json.Marshal(qa.Steps)
	_, err := db.Exec("INSERT INTO questions (topic, subtopic, question, answer, steps) VALUES ($1, $2, $3, $4, $5)", qa.Topic, qa.Subtopic, qa.Question, qa.Answer, string(steps))
	return err
}

func updateQA(qa QA) error {
	steps, _ := json.Marshal(qa.Steps)
	_, err := db.Exec("UPDATE questions SET topic = $1, subtopic = $2, question = $3, answer = $4, steps = $5 WHERE id = $6", qa.Topic, qa.Subtopic, qa.Question, qa.Answer, string(steps), qa.ID)
	return err
}

func deleteQA(id int) error {
	_, err := db.Exec("DELETE FROM questions WHERE id = $1", id)
	return err
}

func readCheckmarks() (map[int][]int, error) {
	data, err := os.ReadFile("checkmarks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[int][]int), nil
		}
		return nil, err
	}
	var checkmarks map[int][]int
	json.Unmarshal(data, &checkmarks)
	return checkmarks, nil
}

func resetCompletionForID(checkmarks map[int][]int, id int) {
	if checkmarks[id] != nil {
		checkmarks[id] = []int{}
		err := updateCheckmarks(checkmarks)
		if err != nil {
			log.Fatal(err)
		}
		color.New(color.FgGreen).Printf("Checkmarks reset successfully for ID %d.\n", id)
	} else {
		color.New(color.FgRed).Printf("No checkmarks found for ID %d.\n", id)
	}
}

func resetAllCheckmarks(checkmarks map[int][]int) {
	for id := range checkmarks {
		checkmarks[id] = []int{}
	}
	err := updateCheckmarks(checkmarks)
	if err != nil {
		log.Fatal(err)
	}
	color.New(color.FgGreen).Println("All checkmarks reset successfully.")
}
func getQAsByID(ids []int) ([]QA, error) {
	allQA, err := listAll()
	if err != nil {
		return nil, err
	}

	var specificQAs []QA
	for _, id := range ids {
		for _, qa := range allQA {
			if qa.ID == id {
				specificQAs = append(specificQAs, qa)
				break
			}
		}
	}

	return specificQAs, nil
}

func updateCheckmarks(checkmarks map[int][]int) error {
	data, err := json.Marshal(checkmarks)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("checkmarks.json", data, 0644)
}

func main() {
	printLogo()
	color.New(color.FgCyan).Printf("Welcome to StepQA!\n\n")
	checkmarks, err := readCheckmarks()
	if err != nil {
		log.Fatal(err)
	}

	for {
		actions := []string{
			"List QAs by ID",
			"List all Q&A",
			"Search for Q&A",
			"Add Q&A",
			"Update Q&A",
			"Delete Q&A",
			"Toggle step completion",
			"Reset all step checkmarks",
			"Reset step checkmarks for a specific ID",
			"Exit",
		}

		// Calculate the number of lines required to display the longest action string
		maxLines := 0
		for _, action := range actions {
			lines := strings.Count(action, "\n") + 1
			if lines > maxLines {
				maxLines = lines
			}
		}

		actionPrompt := promptui.Select{
			Label: "Choose an action",
			Items: actions,
			Size:  len(actions) * maxLines, // Set the Size property according to the calculated maxLines
		}

		actionIndex, _, err := actionPrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
			return
		}

		switch actions[actionIndex] {
		case "List QAs by ID":
			idPrompt := promptui.Prompt{
				Label: "Enter the IDs of the Q&As to display (comma-separated)",
			}
			idStr, err := idPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed: %v\n", err)
				return
			}
			idStrs := strings.Split(idStr, ",")
			ids := make([]int, len(idStrs))
			for i, idStr := range idStrs {
				ids[i], _ = strconv.Atoi(strings.TrimSpace(idStr))
			}
			specificQAs, err := getQAsByID(ids)
			if err != nil {
				log.Fatal(err)
			}
			displayQAList(specificQAs, checkmarks)
		case "List all Q&A":
			allQA, err := listAll()
			if err != nil {
				log.Fatal(err)
			}
			displayQAList(allQA, checkmarks)
		case "Search for Q&A":
			searchPrompt := promptui.Prompt{
				Label: "Search",
			}
			searchTerm, err := searchPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed: %v\n", err)
				return
			}
			results, err := searchQA(searchTerm)
			if err != nil {
				log.Fatal(err)
			}
			displayQAList(results, checkmarks)
		case "Add Q&A":
			qa := inputQAInfo(QA{})
			err = addQA(qa)
			if err != nil {
				log.Fatal(err)
			}
			color.New(color.FgGreen).Printf("Q&A added successfully.\n")
		case "Update Q&A":
			updateIDPrompt := promptui.Prompt{
				Label: "Enter the ID of the Q&A to update",
			}
			updateIDStr, err := updateIDPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed: %v\n", err)
				return
			}
			fmt.Println()
			updateID, _ := strconv.Atoi(updateIDStr)
			allQA, err := listAll()
			if err != nil {
				log.Fatal(err)
			}
			var updateQAEntry QA
			for _, qa := range allQA {
				if qa.ID == updateID {
					updateQAEntry = qa
					break
				}
			}
			if updateQAEntry.ID == 0 {
				color.New(color.FgRed).Printf("No Q&A found with the provided ID.\n")
			} else {
				updatedQA := inputQAInfo(updateQAEntry)
				err = updateQA(updatedQA)
				if err != nil {
					log.Fatal(err)
				}
				color.New(color.FgGreen).Printf("Q&A updated successfully.\n")
			}
		case "Delete Q&A":
			deleteIDPrompt := promptui.Prompt{
				Label: "Enter the ID of the Q&A to delete",
			}
			deleteIDStr, err := deleteIDPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed: %v\n", err)
				return
			}
			deleteID, _ := strconv.Atoi(deleteIDStr)
			err = deleteQA(deleteID)
			if err != nil {
				log.Fatal(err)
			}
			color.New(color.FgGreen).Printf("Q&A deleted successfully.\n")
		case "Toggle step completion":
			toggleCompletion(checkmarks)
			err = updateCheckmarks(checkmarks)
			if err != nil {
				log.Fatal(err)
			}
		case "Reset all step checkmarks":
			resetAllCheckmarks(checkmarks)
		case "Reset step checkmarks for a specific ID":
			resetIDPrompt := promptui.Prompt{
				Label: "Enter the ID of the Q&A to reset checkmarks",
			}
			resetIDStr, err := resetIDPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed: %v\n", err)
				return
			}
			resetID, _ := strconv.Atoi(resetIDStr)
			resetCompletionForID(checkmarks, resetID)
		case "Exit":
			color.New(color.FgCyan).Printf("Goodbye!\n")
			return
		}
	}
}
