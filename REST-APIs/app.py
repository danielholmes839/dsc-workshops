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
    return jsonify(users)


@app.route("/user/<int:user_id>", methods=['GET'])
def get_user(user_id: int):
    """ Endpoint to get a specific user """
    try:
        return users[user_id]
    except:
        return "null", 404


@app.route("/user", methods=['POST'])
def post_user(user):
    """ Endpoint to get a specific user """
    user = request.json
    users.append(user)
    return user

@app.route("/user/<int:user_id>", methods=['GET'])
def put_user():
    pass

@app.route("/user/<int:user_id>", methods=['GET'])
def delete_user():
    pass


app.run(port=3000)