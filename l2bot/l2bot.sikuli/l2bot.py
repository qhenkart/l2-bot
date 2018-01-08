import lib
import dungeons
import weeklys

from org.sikuli.natives import Vision

Vision.setParameter("MinTargetSize", 5)
setAutoWaitTimeout(2.0)
 
todos = {"temple": True, "summoning": True, "sub": True, "daily": True}


def main():
    global todos
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        lib.login()
        lib.checkOK()
        lib.checkSkip()
        lib.reset()

        if todos["temple"] or todos["summoning"]:
            todos = dungeons.run(todos)
            lib.reset()
            if exists("1515159000793.png"):
                click("1515159012057.png")

        if todos["sub"]:
            if exists("1515079317950.png"):
                    print("claiming reward")
                    click("1515079317950.png")
        
            todos = weeklys.sub(todos)

        if todos["daily"]:
            lib.reset()
            if exists("1514888901019-1.png"):
                click("1514888901019-1.png")
            if exists("1515047937858.png"):
                click("1515047937858.png")
                todos = acceptDailys(todos)

        if not todos["sub"] and not todos["temple"] and not todos["summoning"]:
            lib.reset()
            todos = weeklys.run(todos)


while True:
    try:
        print("working")
        main()
    except:
        print("todos: ", todos)
        print("failed running again")
        main()