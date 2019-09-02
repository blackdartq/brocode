#!/usr/bin/python3

import sys

def get_navbar_section(file_name):
    navbar_to_transfer = open(file_name, "r")
    sections = "".join(navbar_to_transfer.readlines()).split("<navbar>")
    navbar_and_end_sections = sections[1].split("</navbar>")
    navbar_section = "              <navbar>\n" + navbar_and_end_sections[0] + "\n          </navbar>"
    return navbar_section

def replace_navbar_section(file_name, navbar_section):
    navbar_being_replaced = open(file_name, "r")
    sections = "".join(navbar_being_replaced.readlines()).split("<navbar>")
    beginning_section = sections[0]
    if len(sections) != 2:
        print("error in the middle to end sections {}".format(len(sections)))
        return
    navbar_and_end_sections = sections[1].split("</navbar>")
    end_section = navbar_and_end_sections[1]
    text = "{}\n{}\n{}".format(beginning_section, navbar_section, end_section)
    navbar_being_replaced.close()
    navbar_being_replaced = open(file_name, "w")
    navbar_being_replaced.write(text)



if len(sys.argv) != 3:
    print("please give two files as arguments!")
else:    
    # File names from the arguements
    file1 = sys.argv[1]
    file2 = sys.argv[2]

    navbar_section_file_1 = get_navbar_section(file1)
    print(navbar_section_file_1)
    if input("do you want to replace navbar with this one?(y/n): ") == "y":
        print("moving navbar from {} to {}".format(file1, file2))
        replace_navbar_section(file2, navbar_section_file_1)
    else:
        print("canceling!!!")
    





