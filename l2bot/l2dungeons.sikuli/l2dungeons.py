import lib
import dungeons

from org.sikuli.natives import Vision

Vision.setParameter("MinTargetSize", 5)
setAutoWaitTimeout(2.0)
 
todos = {"temple": True, "summoning": True}


def main():
    global todos
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        lib.checkOK()
        lib.checkSkip()
        lib.reset()
        todos = dungeons.run(todos)



lib.login()
while True:
    if not todos["temple"] and not todos["summoning"]:
        break
    try:
        print("working")
        main()
    except:
        print("failed running again")
        lib.login()
        main()