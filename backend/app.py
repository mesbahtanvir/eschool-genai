from flask import Flask, request, jsonify
from flask_cors import CORS
from openai import OpenAI
import os
import json


# Initialize the Flask app
app = Flask(__name__)

# Enable CORS for all routes and methods
CORS(app)

client = OpenAI(api_key=os.getenv("OPENAI_API_KEY"))


course_blueprint = {
    "title" : "A very intuitive title of the course",
    "overview" : "A very good overview",
    "objective": "A very achieveable objective",
    "modules": [
        {
            "title": "Introduction to the course",
            "topcis": ["overview", "..."]
        }
    ]
}

def generate_text(prompt, model="gpt-3.5-turbo", max_tokens=50):
    """
    Helper function to call OpenAI's API to generate text based on a given prompt.
    """
    try:
        completion = client.chat.completions.create(
            messages=[
                {
                    "role": "user",
                    "content": prompt,
                }
             ],
            model="gpt-3.5-turbo-1106",
        )
        # Extract the assistant's reply
        return completion.choices[0].message.content
    except Exception as e:
        print(f"Error calling OpenAI API: {e}")
        return f"Error calling OpenAI API: {e}"

# Define an endpoint to accept a text input via query parameter and return the generated output
@app.route('/generate', methods=['GET', 'OPTIONS'])
def generate():
    if request.method == 'OPTIONS':
        # Properly respond to the preflight request
        response = app.make_default_options_response()
        return response
    # Get the input_text from the query parameters
    input_text = request.args.get('input_text')
    input_text = "give me a course outline for the topic " + input_text + \
        "given that i dont have any prior knowledge on this topic" +  \
        "Give me the response as json object and take this as a blueprint," + \
        json.dumps(course_blueprint)

    print(input_text)
    
    # Check if the input text is provided
    if not input_text:
        return jsonify({"error": "No input text provided"}), 400
    
    # Call the OpenAI API to generate text
    output_text = generate_text(input_text)
    
    if output_text is None:
        return jsonify({"error": "Failed to get a response from OpenAI API"}), 500
    output = json.loads(output_text)
    # Return the generated text as a JSON response
    return jsonify({"input": input_text, "output": output})

# Run the app
if __name__ == "__main__":
    app.run(debug=True)
