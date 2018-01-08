import lib
import weeklys

from org.sikuli.natives import Vision

#Vision.setParameter("MinTargetSize", 5)

setAutoWaitTimeout(2.0)

todos = {"daily": True, "weekly": True, "sub": True}

def main(): 
    global todos
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        lib.checkOK()
        lib.checkSkip()
        lib.reset()
        if exists("1515079317950.png"):
            click("1515079317950.png")

        if not todos["weekly"] and todos["sub"]:
            print("running sub")
            todos = weeklys.sub(todos)
        else:
            print("running weeklys")
            todos = weeklys.run(todos)


lib.login()
while True:
    if not todos["weekly"] and not todos["sub"]:
        print("finished with script")
        break
    try:
        print("working")
        main()
    except:
        print("failed running again")
        lib.login()
        main()