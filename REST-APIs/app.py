from flask import Flask, request, jsonify

app = Flask(__name__)


next_id = 1
users = {
    0: {
        'name': 'Daniel',
        'age': 20,
    }
}

@app.route("/users", methods=['GET'])
def get_users():
    """ show all users """
    return jsonify([{'id': key, **user} for key, user in users.items()])


@app.route("/user/<int:user_id>", methods=['GET'])
def get_user(user_id: int):
    """ Endpoint to get a specific user """
    if user_id in users:
        return users[user_id]
    
    return "NOT FOUND", 404

@app.route("/user", methods=['POST'])
def post_user():
    """ Endpoint to get a specific user """
    user = request.json
    users.append(user)
    return "OK", 200

@app.route("/user/<int:user_id>", methods=['GET'])
def put_user(user_id: int):
    users[user_id] = request.json
    return "OK", 200

@app.route("/user/<int:user_id>", methods=['GET'])
def delete_user(user_id: int):
    users.pop(user_id)
    return "OK", 200
    

app.run(port=3000)