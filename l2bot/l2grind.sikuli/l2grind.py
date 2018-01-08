import lib



from org.sikuli.natives import Vision

Vision.setParameter("MinTargetSize", 9)

def main():
    nox = App.open("Nox App Player")
    with Region(nox.window()):

        lib.checkOK()
        lib.reset()
        if not exists("1514949313948.png"):
            print("looking for auto")
            if exists("1514955626442.png"):
                print("hitting auto")
                click("1514955626442.png")

lib.login()
while True:
    try:
        print("working")
        main()
    except:
        print("failed running again")
        lib.login()
        main()