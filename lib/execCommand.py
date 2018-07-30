import subprocess, json, platform

class Node:
    def __init__(self):
        self.distro = ""
        self.model = ""
        self.kernel = ""
        self.cpu_usage = ""
        self.cpu_temp = ""
        self.ram_usage = ""
        self.swap_usage = ""
        self.hdd_usage = ""
        self.storage_usage = ""

n = Node()

if platform.system() == "Linux":
    cmd = "uname -r | head --bytes -1"
    n.kernel = subprocess.check_output(cmd, shell = True )

    cmd = "cat /etc/os-release | head -1 | cut -d '\"' -f 2 | head --bytes -1"
    n.distro = subprocess.check_output(cmd, shell = True )

    cmd = "dmesg | grep Machine | cut -d ':' -f 4 | head --bytes -1"
    n.model = subprocess.check_output(cmd, shell = True )

    cmd = "top -bn1 | grep load | awk '{printf \"%.2f\", $(NF-2)}'"
    n.cpu_usage = subprocess.check_output(cmd, shell = True )

    cmd = "vcgencmd measure_temp | cut -d '=' -f 2 | head --bytes -3"
    n.cpu_temp = subprocess.check_output(cmd, shell = True )

    cmd = "free -m | awk 'NR==2{printf \"%.2f\", $3*100/$2 }'"
    n.ram_usage = subprocess.check_output(cmd, shell = True )

    cmd = "free -m | awk 'NR==3{printf \"%.2f\", $3*100/$2 }'"
    n.swap_usage = subprocess.check_output(cmd, shell = True )

    cmd = "df -h | awk '$NF==\"/\"{printf \"%s\", $5}' | head --bytes -1"
    n.hdd_usage = subprocess.check_output(cmd, shell = True )

    cmd = "df -h | awk '$NF==\"/hydra/storage\"{printf \"%s\", $5}' | head --bytes -1"
    n.storage_usage = subprocess.check_output(cmd, shell = True )

print(json.dumps(n.__dict__))
