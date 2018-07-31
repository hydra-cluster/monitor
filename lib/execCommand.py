import subprocess, json, platform

class Param:
    def __init__(self, value, label, unit):
        self.Value = value
        self.Label = label
        self.Unit = unit

class Node:
    def __init__(self):
        self.Distro = ""
        self.Model = ""
        self.Kernel = ""
        self.Params = []

    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__, sort_keys=True, indent=4)

n = Node()

if platform.system() == "Linux":
        
    cmd = "uname -r | head --bytes -1"
    n.kernel = subprocess.check_output(cmd, shell = True )

    cmd = "cat /etc/os-release | head -1 | cut -d '\"' -f 2 | head --bytes -1"
    n.distro = subprocess.check_output(cmd, shell = True )

    cmd = "dmesg | grep Machine | cut -d ':' -f 4 | head --bytes -1"
    n.model = subprocess.check_output(cmd, shell = True )

    cmd = "top -bn1 | grep load | awk '{printf \"%.2f\", $(NF-2)}'"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "CPU Load", ""))

    cmd = "vcgencmd measure_temp | cut -d '=' -f 2 | head --bytes -3"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "CPU Temp", "C"))

    cmd = "free -m | awk 'NR==2{printf \"%.2f\", $3*100/$2 }'"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "RAM", "%"))

    cmd = "free -m | awk 'NR==3{printf \"%.2f\", $3*100/$2 }'"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "SWAP", "%"))
    
    cmd = "df -h | awk '$NF==\"/\"{printf \"%s\", $5}' | head --bytes -1"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "HDD", "%"))
    
    cmd = "df -h | awk '$NF==\"/hydra/storage\"{printf \"%s\", $5}' | head --bytes -1"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "Storage", "%"))

print(n.toJSON())
