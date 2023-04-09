-- Database schema
CREATE TABLE IF NOT EXISTS questions (
                                         id SERIAL PRIMARY KEY,
                                         topic TEXT NOT NULL,
                                         subtopic TEXT NOT NULL,
                                         question TEXT NOT NULL,
                                         answer TEXT NOT NULL,
                                         steps TEXT NOT NULL
);

-- Sample questions with steps
INSERT INTO questions (topic, subtopic, question, answer, steps)
VALUES ('Cooking', 'Baking', 'How do you bake chocolate chip cookies?', 'Follow these steps to bake chocolate chip cookies.', '["1. Preheat the oven to 350°F (180°C).", "2. In a bowl, mix together butter, sugar, and brown sugar.", "3. Add eggs and vanilla extract, and mix well.", "4. Stir in flour, baking soda, and salt.", "5. Fold in chocolate chips.", "6. Drop spoonfuls of dough onto a baking sheet.", "7. Bake for 10-12 minutes or until golden brown."]');

INSERT INTO questions (topic, subtopic, question, answer, steps)
VALUES ('Gardening', 'Plant Care', 'How do I repot a houseplant?', 'Follow these steps to repot a houseplant.', '["1. Choose a new pot that is 1-2 inches larger in diameter than the current pot.", "2. Fill the new pot with a layer of fresh potting soil.", "3. Carefully remove the plant from its current pot.", "4. Place the plant in the new pot and add more potting soil, pressing gently around the roots.", "5. Water the plant thoroughly."]');

INSERT INTO questions (topic, subtopic, question, answer, steps)
VALUES ('Technology', 'Smartphones', 'How do I factory reset an Android phone?', 'Follow these steps to factory reset an Android phone.', '["1. Open the Settings app.", "2. Scroll down and tap on System.", "3. Tap on Reset options.", "4. Select Erase all data (factory reset).", "5. Tap on Reset phone, and enter your PIN or password if prompted.", "6. Confirm the action by tapping on Erase everything."]');

INSERT INTO questions (topic, subtopic, question, answer, steps)
VALUES ('Finance', 'Budgeting', 'How do I create a personal budget?', 'Follow these steps to create a personal budget.', '["1. Gather all your financial information, such as bank statements and bills.", "2. Calculate your total monthly income.", "3. List all your monthly expenses.", "4. Divide your expenses into fixed and variable categories.", "5. Set spending limits for each expense category.", "6. Track your spending and adjust your budget as needed."]');

INSERT INTO questions (topic, subtopic, question, answer, steps)
VALUES ('Health', 'Exercise', 'How do I start a workout routine?', 'Follow these steps to start a workout routine.', '["1. Set realistic fitness goals.", "2. Choose a workout program that fits your schedule and preferences.", "3. Invest in necessary equipment or a gym membership.", "4. Warm up and stretch before each workout.", "5. Follow your workout plan consistently.", "6. Gradually increase intensity and duration as you progress."]');

INSERT INTO questions (topic, subtopic, question, answer, steps)
VALUES ('Travel', 'Packing', 'What should I pack for a week-long trip?', 'Follow these steps to pack for a week-long trip.', '["1. Check the weather forecast for your destination.", "2. Make a packing list, including clothing, toiletries, and electronics.", "3. Choose a suitcase or backpack that meets airline requirements.", "4. Pack versatile clothing items that can be mixed and matched.", "5. Roll or fold clothes to save space.", "6. Pack travel-sized toiletries and essentials in a separate bag.", "7. Do not forget important documents, such as your passport and travel insurance."]');

INSERT INTO questions (topic, subtopic, question, answer, steps)
VALUES ('Education', 'Study Tips', 'How do I prepare for an exam?', 'Follow these steps to prepare for an exam.', '["1. Review your course materials and identify important topics.", "2. Create a study schedule with specific goals and deadlines.", "3. Break down complex topics into smaller sections.", "4. Use active learning techniques, such as flashcards and practice tests.", "5. Review your notes and textbook regularly.", "6. Ask for help if needed, and join study groups.", "7. Get plenty of rest and maintain a healthy lifestyle."]');

INSERT INTO questions (topic, subtopic, question, answer, steps)
VALUES ('Automotive', 'Car Maintenance', 'How do I change a flat tire?', 'Follow these steps to change a flat tire.', '["1. Find a safe location to park your car and turn on your hazard lights.", "2. Apply the parking brake and place wheel chocks or rocks behind the tires.", "3. Retrieve your spare tire, jack, and lug wrench from the trunk.", "4. Loosen the lug nuts before lifting the car.", "5. Place the jack under the car and raise it until the flat tire is off the ground.", "6. Remove the lug nuts and the flat tire, and replace it with the spare tire.", "7. Tighten the lug nuts and lower the car."]');

INSERT INTO questions (topic, subtopic, question, answer, steps)
VALUES ('Home Improvement', 'Painting', 'How do I paint a room?', 'Follow these steps to paint a room.', '["1. Choose your paint color and finish.", "2. Clear the room of furniture and cover the floor with drop cloths.", "3. Clean and repair the walls, filling any holes or cracks.", "4. Apply painters tape around windows, doors, and trim.", "5. Prime the walls, if necessary.", "6. Start painting the edges with a brush, then fill in the walls using a roller.", "7. Apply multiple coats if needed, allowing each coat to dry before applying the next."]');

INSERT INTO questions (topic, subtopic, question, answer, steps)
VALUES ('Pets', 'Dog Training', 'How do I teach my dog to sit?', 'Follow these steps to teach your dog to sit.', '["1. Grab a treat and hold it close to your dogs nose.", "2. Slowly move the treat up and back, guiding your dogs head to follow.", "3. As your dogs head moves up, their bottom should lower to the ground.", "4. When your dogs bottom touches the ground, say "Sit" and reward them with the treat.", "5. Practice this command regularly and gradually phase out the treat."]');