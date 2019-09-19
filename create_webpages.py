#!/usr/bin/python3
import os

pug_dir = "./web/pug/"
sass_dir = "./web/sass/"
web_dir = "./web/"

print("Converting pug files to HTML...")
pug_files = os.listdir(pug_dir)
print("Creating html files from {}".format(pug_files))
for pug_file in pug_files:
    if ".pug" in pug_file:
        os.system("pug -P -o {} {}{} ".format(web_dir, pug_dir, pug_file))

# Converting sass files to css
print("Converting SASS files to CSS...")
sass_files = os.listdir(sass_dir)
print("Creating css files from {}".format(sass_files))
for sass_file in sass_files:
    if ".sass" in sass_file:
        if "_" not in sass_file:
            # removes the sass from filename and adds css to the end
            css_file = sass_file[0:len(sass_file)-4] + "css"
            os.system("sass {}{} {}{} ".format(sass_dir, sass_file, web_dir, css_file))
