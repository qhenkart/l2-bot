from sikuli import *

import lib

quest = Region(49,266,272,112)

reward = Region(425,631,530,75)


cleared = Region(1025,658,199,66)

def runQuest():
    c = 1
    while not reward.exists("1514880563893.png"):
        print("in loop" + str(c))
        c = c + 1
        lib.checkOK()
        lib.checkSkip()
        if quest.exists("1514882132041.png"):
            print("clicking to accept quest")
            quest.click("1514867316210.png")
            lib.checkSkip()
            
        lib.checkWalk()
        print("checking accept")
        if exists("1514880148209.png"):
            click("1514880160477.png")
            print("accepting quest")
        lib.checkSkip()
        
    return;
        

def run():
    if quest.exists("1514882132041.png"):
        print("clicking to accept quest")
        quest.click("1514867316210.png")
        lib.checkSkip()
        #loops until claim reward is seen
        runQuest()
    
    #finish quest           
    if reward.exists("1514880563893.png"):
        print("claiming reward")
        reward.click("1514880575821.png")
    #story progression
    if cleared.exists("1514886419623.png"):
        print("story complete")
        cleared.click("1514886419623.png")
                        
    return;
    



  