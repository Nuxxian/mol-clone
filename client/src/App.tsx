// src/components/QuestionnaireForm.tsx

import { createSignal } from 'solid-js';

export default function App() {
  const [answers, setAnswers] = createSignal({
    question1: '',
    question2: '',
  });

  const handleInputChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    setAnswers({
      ...answers(),
      [target.name]: target.value,
    });
  };

  const handleSubmit = (event: Event) => {
    event.preventDefault();
    console.log(event)
    console.log('Form submitted with answers:', answers());
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Question 1: What is your favorite color?</label>
        <div>
          <label>
            <input
              type="radio"
              name="question1"
              value="Red"
              checked={answers().question1 === 'Red'}
              onChange={handleInputChange}
            />
            Red
          </label>
          <label>
            <input
              type="radio"
              name="question1"
              value="Blue"
              checked={answers().question1 === 'Blue'}
              onChange={handleInputChange}
            />
            Blue
          </label>
          <label>
            <input
              type="radio"
              name="question1"
              value="Green"
              checked={answers().question1 === 'Green'}
              onChange={handleInputChange}
            />
            Green
          </label>
        </div>
      </div>

      <div>
        <label>Question 2: Do you like coding?</label>
        <div>
          <label>
            <input
              type="radio"
              name="question2"
              value="Yes"
              checked={answers().question2 === 'Yes'}
              onChange={handleInputChange}
            />
            Yes
          </label>
          <label>
            <input
              type="radio"
              name="question2"
              value="No"
              checked={answers().question2 === 'No'}
              onChange={handleInputChange}
            />
            No
          </label>
        </div>
      </div>

      <button type="submit">Submit</button>
    </form>
  );
};

