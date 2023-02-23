import subprocess

# command and arguments
command = "python"
args = [
    "demo/restoration_video_demo.py", 
    "./configs/restorers/basicvsr/basicvsr_reds4.py", 
    "https://download.openmmlab.com/mmediting/restorers/basicvsr/basicvsr_reds4_20120409-0e599677.pth", 
    "data/Vid4/BIx4/calendar/", 
    "./output"
]

process = subprocess.Popen([command, *args], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
output, error = process.communicate()
