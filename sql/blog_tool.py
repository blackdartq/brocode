#!/usr/bin/python3
import os;

print("Welcome to the blog tool!")
user_input = input("1. see all posts\n2. add blog post\nWhat do you want to do?: ")
if user_input == "1":
    database_query = "psql -U blackdartq -d brocode -c \"SELECT * FROM blog;\""
    os.system(database_query)
elif user_input == "2":
    name = input("what is the name of your post?: ")
    post = input("type post below\n")
    database_query = "psql -U blackdartq -d brocode -c \"INSERT INTO blog VALUES(DEFAULT, '{}', '{}')\";".format(name, post)
    os.system(database_query)
else:
    print("please type one of the options!")
