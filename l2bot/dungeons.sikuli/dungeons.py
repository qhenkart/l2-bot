from sikuli import *
import lib

def run(todos):
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        if not exists("1515140799080.png",0) and not exists("1515142056473.png",0):
            return todos
        if exists("1515140799080.png",0):
            print("clicking party button")
            click("1515140799080.png")
        if exists("1515140814786.png"):
            print("search for party")
            click("1515140814786.png")
        if exists("1515140834248.png"):
            print("auto join")
            click("1515140841821.png")
        if exists("1515142089421.png"):
            print("auto join from autojoining party menu")
            click("1515142097928.png")
        if exists("1515142048219.png"):
            print("auto join from auto join screen")
            click("1515142056473.png")
        if todos['temple']:
            if exists("1515140878528.png"):
                print("temple")
                click("1515140885522.png")
                wait(2)
                if exists("1515140907216.png"):
                    click("1515140907216.png")
                    click("1515140943771.png")
                    click("1515140956348.png")
                    if exists("1515144177462.png"):
                        todos['temple'] = False
                        lib.checkOK()
        if todos['summoning']:          
            if exists("1515141671416.png"):
                print("summoning")
                click("1515141671416.png")
                wait(2)
                if exists("1515141688796.png"):
                    click("1515141688796.png")
                    click("1515141704203.png")
                    click("1515140956348.png")
                    if exists("1515144177462.png"):
                        todos['summoning'] = False
                        lib.checkOK()
    return todos
