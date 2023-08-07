import random as r
import csv
import os
import copy

command = "open /Users/dart/Pictures/Games/RE4Make/RE4.png"

handguns = ["SR-09 R","Punisher","Red9","Blacktail","Matilda","Sentinel Nine"]
shotguns = ["W-870", "Riot Gun", "Striker", "", "Skull Shaker"]
rifles = ["SR M1903", "Stingray", "CQBR Assault Rifle"]
magnums = ["Broken Butterfly", "Killer7"]
specials = ["Bolt Thrower", "Infinite Rocket Launcher", "Chicago Sweeper"]


weaponsList = handguns + shotguns + rifles + magnums + specials
currentGuns = copy.deepcopy(weaponsList)

helpString = "h for help ; s to save ; r to roll ; l to load; g to show guide "
currentChapter = 1
finalChapter = 16
pickedGun = ""
rows = []
def save():
    with open("guns.csv", 'w', encoding='UTF8', newline='') as f:
        writer = csv.writer(f)
        print(rows)
        writer.writerows(rows)
    return

def load():
    loadedRows = []
    global currentChapter, rows, currentGuns
    loadedGuns = copy.deepcopy(weaponsList)
    with open('guns.csv') as csv_file:
        csv_reader = csv.reader(csv_file, delimiter=',')
        for row in csv_reader:
            chapter = int(row[0])
            #print(chapter)
            pickedGun = row[1]
            #print(pickedGun)
            loadedGuns.remove(pickedGun)
            loadedRows.append([chapter, pickedGun])
    #print(loadedRows)
    rows = copy.deepcopy(loadedRows)
    currentGuns = copy.deepcopy(loadedGuns)
    currentChapter = chapter+1
    print("Successfully Loaded Chapter:%d, Using:%s" % (chapter, pickedGun) )
    return
    
def roll():
    global pickedGun
    global currentChapter
    pickedGun = r.choice(currentGuns)
    rows.append([currentChapter, pickedGun])
    currentGuns.remove(pickedGun)
    print("Chapter %d, Using the %s" % (currentChapter, pickedGun))
    currentChapter += 1
    return
    
def switch(lang):
    if lang == 'h':
        print(helpString)
    elif lang == 's':   
        save()
    elif lang == 'r':
        roll()
    elif lang == 'l':
        load()
    elif lang == 'g':
        os.system(command)
    elif lang == 'i':
        print(rows)
    else:
        print("Invalid input, try again.")

print("Welcome to the RE4 Make Pick-A-Gun Service Version 2!")

while currentChapter < finalChapter:
    print("Currently in Chapter %s" % currentChapter)
    choice = input("What are ya buyin: ") # save this input later
    switch(choice)
exit(0)
