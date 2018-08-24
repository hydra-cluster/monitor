import subprocess, json, platform, argparse

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

parser = argparse.ArgumentParser()
parser.add_argument("-cmd", help="Define a command to be executed if not defined return agent system params")
args = parser.parse_args()

if args.cmd:
    try:
        output = subprocess.check_output(args.cmd, stderr=subprocess.STDOUT, shell = True)
        result = {'status': 'done', 'output': output}
    except subprocess.CalledProcessError as e:
        result = {'status': 'error', 'output': e.output}
    print(json.dumps(result, default=lambda o: o.__dict__, sort_keys=True, indent=4))
else: 
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

        cmd = "cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq"
        value = subprocess.check_output(cmd, shell = True )
        n.Params.append(Param(value, "CPU Freq", "Hz", "", "" ))

        cmd = "free -m | awk 'NR==2{printf \"%.2f\", $2 }'"
        value = subprocess.check_output(cmd, shell = True )
        n.Params.append(Param(value, "RAM Total", "MB", "70", "85" ))

        cmd = "free -m | awk 'NR==2{printf \"%.2f\", $3 }'"
        value = subprocess.check_output(cmd, shell = True )
        n.Params.append(Param(value, "RAM Used", "MB", "70", "85" ))

        cmd = "free -m | awk 'NR==3{printf \"%s\", $2 }'"
        value = subprocess.check_output(cmd, shell = True )
        n.Params.append(Param(value, "SWAP Total", "MB", "50", "80" ))

        cmd = "free -m | awk 'NR==3{printf \"%s\", $3 }'"
        value = subprocess.check_output(cmd, shell = True )
        n.Params.append(Param(value, "SWAP Used", "MB", "50", "80" ))
        
        cmd = "df | awk '$NF==\"/\"{printf \"%s\", $2}'"
        value = subprocess.check_output(cmd, shell = True )
        n.Params.append(Param(value, "SD Card Total", "Bytes", "50", "85" ))

        cmd = "df | awk '$NF==\"/\"{printf \"%s\", $3}'"
        value = subprocess.check_output(cmd, shell = True )
        n.Params.append(Param(value, "SD Card Used", "Bytes", "50", "85" ))
        
        cmd = "df | awk '$NF==\"/hydra/storage\"{printf \"%s\", $2}'"
        value = subprocess.check_output(cmd, shell = True )
        n.Params.append(Param(value, "Storage Total", "Bytes", "70", "85" ))

        cmd = "df | awk '$NF==\"/hydra/storage\"{printf \"%s\", $3}'"
        value = subprocess.check_output(cmd, shell = True )
        n.Params.append(Param(value, "Storage Used", "Bytes", "70", "85" ))
    
    else :
        n.Kernel = "Testing..."
        n.Distro = "Testing..."
        n.Model = "Testing..."
        n.Params.append(Param("0", "CPU Load", "", "0.7", "0.9" ))
        n.Params.append(Param("0", "CPU Temp", "C", "70", "80" ))
        n.Params.append(Param("0", "CPU Freq", "Hz", "", "" ))
        n.Params.append(Param("0", "RAM Total", "MB", "70", "85" ))
        n.Params.append(Param("0", "RAM Used", "MB", "70", "85" ))
        n.Params.append(Param("0", "SWAP Total", "MB", "50", "80" ))
        n.Params.append(Param("0", "SWAP Used", "MB", "50", "80" ))
        n.Params.append(Param("0", "SD Card Total", "Bytes", "50", "85" ))
        n.Params.append(Param("0", "SD Card Used", "Bytes", "50", "85" ))
        n.Params.append(Param("0", "Storage Total", "Bytes", "70", "85" ))
        n.Params.append(Param("0", "Storage Used", "Bytes", "70", "85" ))

    print(n.toJSON())
    