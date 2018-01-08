import quests
import lib
import weeklys


from org.sikuli.natives import Vision

Vision.setParameter("MinTargetSize", 3)

 

nox = App.open("Nox App Player")

while True:   
    with Region(nox.window()):
        lib.login()
        lib.checkOK()
        lib.checkSkip()
        quests.run()