from sikuli import *

import lib

setAutoWaitTimeout(2.0)

def run(todos):
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        if exists("1515319729876.png") and not exists("1515317424466.png"):
            todos["weekly"] = False
            return todos
        print("running weeklys")
        click("1514888901019.png")
        wait(5)
        if todos["daily"] and exists("1515153619450.png"):
            click("1515047937858.png")
            todos = acceptDailys(todos)
            
        click("1514888961518.png")
        wait(5)
        if exists("1515076045929.png"):
            print("quest complete")
            click("1515076058366.png")
            wait(2)
        r = Region(308,213,1107,678)
        if r.exists("1515076074001.png"):
            print("Start quest")
            r.click("1515076087648.png")
            wait(2)
            
        if exists("1514889020863.png"):
           click("1514889029232.png")
        wait(2)
        lib.checkWalk()
    
    
        return todos;

def acceptDailys(todos):
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        print("checking dailys")
    
        r = Region(310,298,1069,558)
        while r.exists("1515047471207.png"):
            print("claiming dailys reward")
            r.click("1515047486964.png")
            wait(2)
        while r.exists("1515047512032.png"):
            if exists("1515048738582.png"):
                print("finished with quests")
                lib.checkOK()
                return;
            print("starting quest")
            r.click("1515047522216.png")
            wait(2)
            if exists("1515154167562.png"):
                todos["daily"] = False
                print("finished todos")
                lib.checkOK()
                return todos
        return todos

def sub(todos):
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        if exists("1515319729876.png") and not exists("1515156791885.png"):
            print("finished with subs")
            todos["sub"] = False
            return todos
        if exists("1515078806613.png"):
            print("clicking quest available")
            click("1515078819634.png")
        if exists("1515078870435.png"):
            print("clicking fulfill")
            click("1515078889568.png")
            lib.checkOK()
            wait(3)
        if exists("1515078927940.png"):
            print("starting quest")
            click("1515078935414.png")
            lib.checkWalk()
    return todos