import lib
import weeklys

from org.sikuli.natives import Vision

Vision.setParameter("MinTargetSize", 5)

 
setAutoWaitTimeout(2.0)

todos = {"sub": True}

def main(): 
    global todos
    nox = App.open("Nox App Player")
    with Region(nox.window()):
        if todos["sub"]:

            lib.checkOK()
            lib.checkSkip()
            lib.reset()
            if exists("1515079317950.png"):
                print("claiming reward")
                click("1515079317950.png")
    
            todos = weeklys.sub(todos)


lib.login()
while True:
    if not todos["sub"]:
        break
    try:
        print("working")
        main()
    except:
        print("failed running again")
        lib.login()
        main()