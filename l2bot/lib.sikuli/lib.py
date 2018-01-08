from sikuli import *

def reset():
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        if exists("1515048421156.png"):
            click("1515048428134.png")
        return;
    
def checkSkip():
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        print("checking skip")
        if exists("1514890247176.png"):
            print("skip exists, clicking") 
            click("1514890255615.png")
        return;

def checkOK():
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        r = Region(216,155,1149,723)
        print("check ok")
        if r.exists("1514904011028.png"):
            print("hitting ok")
            r.click("1514904019510.png")
        if r.exists("1514904513357.png"):
            print("connect now")
            r.click("1514904513357-1.png")
    
        return;

def checkWalk(): 
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        print("checking walk")
        if exists("1514879730198.png"):
            print("walk found")
            click("1514879742678.png")
        return;

def login():
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        if exists("1515302151648.png", 0):
            print("opening application")
            click("1515302158596.png")
            wait(40)
        print("checking login")
        if exists("1514904310014.png"):
            print("tap to start")
            click("1514904310014.png")
            wait(15)
        if exists("1515302239813.png"):
            print("x out of notice")
            click("1515302261645.png")
            wait(15)
        if exists("1515302295797.png"):
            print("play")
            click("1514904411742.png")
            wait(30)
        if exists("1515302239813.png"):
            print("x out of notice")
            click("1515302261645.png")
            wait(10)
        