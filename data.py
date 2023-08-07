import random as r
import csv
import os
import copy

class Gun:
    def __init__(self):
        self.command = "open /Users/dart/Pictures/Games/RE4Make/RE4.png"
        handguns = ["SR-09 R","Punisher","Red9","Blacktail","Matilda","Sentinel Nine"]
        shotguns = ["W-870", "Riot Gun", "Striker", "Skull Shaker"]
        rifles = ["SR M1903", "Stingray", "CQBR Assault Rifle"]
        magnums = ["Broken Butterfly", "Killer7"]
        specials = ["Infinite Rocket Launcher", "Chicago Sweeper"]
        # Excluded "Bolt Thrower"

        self.helpString = "h for help ; s to save ; r to rolls   ; l to load; g to show guide "
        self.currentChapter = 1         
        self.finalChapter = 16
        self.pickedGun = ""
        self.rows = []
        
        self.weaponsList = handguns + shotguns + rifles + magnums + specials
        print (self.weaponsList)
        self.currentGuns = copy.deepcopy(self.weaponsList)

    def save(self):
        with open("guns.csv", 'w', encoding='UTF8', newline='') as f:
            writer = csv.writer(f)
            print(self.rows)
            writer.writerows(self.rows)
        return ("Data Successfully Writen. Chapter:%d, Using:%s"  % (self.currentChapter, self.pickedGun))
    
    def load(self):
        loadedRows = []
        loadedGuns = copy.deepcopy(self.weaponsList)
        print (loadedGuns)
        with open('guns.csv') as csv_file:
            csv_reader = csv.reader(csv_file, delimiter=',')
            for row in csv_reader:
                chapter = int(row[0])
                #print(chapter)
                pickedGun = row[1]
                #print(pickedGun)       
                loadedGuns.remove(pickedGun)
                loadedRows.append([chapter, pickedGun])
                self.pickedGun = pickedGun
                self.currentChapter = chapter
        #print(loadedRows)
        self.rows = copy.deepcopy(loadedRows)
        response = ("Successfully Loaded Chapter:%d" % (self.currentChapter))
        return response

    def roll(self):
        self.pickedGun = r.choice(self.currentGuns)
        self.rows.append([self.currentChapter, self.pickedGun])
        self.currentGuns.remove(self.pickedGun) 

        response = "Chapter %d, Using the %s. Get outta my face and play" % (self.currentChapter, self.pickedGun)
        return response
    
    def guide(self):        
        os.system(self.command)
    
# gun = Gun()
    
def switch(lang):
    if lang == 'h':
        print(gun.helpString)
    elif lang == 's':   
        print(gun.save())
    elif lang == 'r':
        print(gun.roll())
    elif lang == 'l':
        print(gun.load())
    # elif lang == 'g':
    #     os.system(command)
    elif lang == 'i':
        print(gun.rows)
    else:
        print("Invalid input, try again.")

# print("Welcome to the RE4 Make Pick-A-Gun Service Version 2!")

# while gun.currentChapter < gun.finalChapter:
#     print("Currently in Chapter %s" % gun.currentChapter)
#     choice = input("What are ya buyin: ") # save this input later
#     switch(choice)
# exit(0)
