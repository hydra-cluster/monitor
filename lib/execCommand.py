import subprocess, json, platform

class Param:
    def __init__(self, value, label, unit, warning, danger):
        self.Value = value
        self.Label = label
        self.Unit = unit
        self.Warning = warning
        self.Danger = danger

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
    n.Kernel = subprocess.check_output(cmd, shell = True )

    cmd = "cat /etc/os-release | head -1 | cut -d '\"' -f 2 | head --bytes -1"
    n.Distro = subprocess.check_output(cmd, shell = True )

    cmd = "dmesg | grep Machine | cut -d ':' -f 4 | head --bytes -1"
    n.Model = subprocess.check_output(cmd, shell = True )

    cmd = "top -bn1 | grep load | awk '{printf \"%.2f\", $(NF-2)}'"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "CPU Load", "", "0.7", "0.9" ))

    cmd = "vcgencmd measure_temp | cut -d '=' -f 2 | head --bytes -3"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "CPU Temp", "C", "70", "80" ))

    cmd = "free -m | awk 'NR==2{printf \"%.2f\", $3*100/$2 }'"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "RAM", "%", "0.7", "0.9" ))

    cmd = "free -m | awk 'NR==3{printf \"%.2f\", $3*100/$2 }'"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "SWAP", "%", "0.5", "0.8" ))
    
    cmd = "df -h | awk '$NF==\"/\"{printf \"%s\", $5}' | head --bytes -1"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "HDD", "%", "0.7", "0.9" ))
    
    cmd = "df -h | awk '$NF==\"/hydra/storage\"{printf \"%s\", $5}' | head --bytes -1"
    value = subprocess.check_output(cmd, shell = True )
    n.Params.append(Param(value, "Storage", "%", "0.7", "0.9" ))

print(n.toJSON())
