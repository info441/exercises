from flask import Flask, jsonify

import os 

app = Flask(__name__)
port = os.environ["PORT"]

@app.route("/v1/image")
def hello():
    return jsonify({ "message": "Hello World from Flask Server!" })

if __name__ == "__main__":
    app.run(debug=True,host='0.0.0.0',port=port)
