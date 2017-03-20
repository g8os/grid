package main

import (
	"encoding/json"
	"net/http"

	"fmt"
	"github.com/g8os/go-client"
	"github.com/gorilla/mux"
)

// Core0API is API implementation of /core0 root endpoint
type Core0API struct {
}

// CoresList is the handler for GET /core0
// List Core0s
func (api Core0API) CoresList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CoreGet is the handler for GET /core0/{id}
func (api Core0API) CoreGet(w http.ResponseWriter, r *http.Request) {
	var respBody Core0
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CoreXList is the handler for GET /core0/{id}/coreX
// List running CoreXses
func (api Core0API) CoreXList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CoreXCreate is the handler for POST /core0/{id}/coreX
// Create a new CoreX
func (api Core0API) CoreXCreate(w http.ResponseWriter, r *http.Request) {
	var reqBody CoreXCreate

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CoreXGet is the handler for GET /core0/{id}/coreX/{coreXid}
// Get CoreX
func (api Core0API) CoreXGet(w http.ResponseWriter, r *http.Request) {
	var respBody CoreX
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CoreXDelete is the handler for DELETE /core0/{id}/coreX/{coreXid}
// Delete CoreX instance
func (api Core0API) CoreXDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CommandsList is the handler for GET /core0/{id}/command
// List running commands
func (api Core0API) CommandsList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CommandGet is the handler for GET /core0/{id}/command/{commandid}
func (api Core0API) CommandGet(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)

	cl := GetConnection(r)

	res, err := cl.Result(client.Job(v["commandid"]), 10)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	var respBody CommandResult
	respBody.Id = res.ID
	respBody.Data = res.Data
	respBody.Level = fmt.Sprintf("%d", res.Level)
	respBody.Name = EnumCommandResultName(res.Command)
	respBody.State = EnumCommandResultState(res.State)
	respBody.Stdout = res.Streams.Stdout()
	respBody.Stderr = res.Streams.Stderr()
	respBody.Starttime = int(res.StartTime)

	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// StateGet is the handler for GET /core0/{id}/core/state
// The aggregated consumption of core0 + all processes (cpu, memory, etc...)
func (api Core0API) StateGet(w http.ResponseWriter, r *http.Request) {
	var respBody CoreStateResult
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// Reboot is the handler for POST /core0/{id}/core/reboot
// Immediately reboot the machine.
func (api Core0API) Reboot(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// Ping is the handler for POST /core0/{id}/core/ping
// Execute a ping command to this Core0
func (api Core0API) Ping(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	cl := GetConnection(r)

	//NOTE: we use Raw instead of the client ping method, because the client ping method
	//is synchronous. Should we change client implementation ?
	job, err := cl.Raw("core.ping", client.A{}, Options(reqBody)...)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	respBody := Location{
		Name: "core.ping",
		Url:  ResultUrl(r, job),
		Id:   string(job),
	}

	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// System is the handler for POST /core0/{id}/core/system
// Execute a new process  on this Core0
func (api Core0API) System(w http.ResponseWriter, r *http.Request) {
	var reqBody CoreSystem

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	cl := GetConnection(r)
	system := client.Core(cl)

	var env map[string]string
	if len(reqBody.Environment) > 0 {
		env = make(map[string]string)
		for _, pair := range reqBody.Environment {
			env[pair.Name] = pair.Value
		}
	}

	job, err := system.SystemArgs(
		reqBody.Name,
		reqBody.Args,
		env,
		reqBody.Pwd,
		reqBody.Stdin,
		Options(reqBody.Command)...)

	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	respBody := Location{
		Name: "core.system",
		Id:   string(job),
		Url:  ResultUrl(r, job),
	}

	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// Kill is the handler for POST /core0/{id}/core/kill
// Kill a process / command
func (api Core0API) Kill(w http.ResponseWriter, r *http.Request) {
	var reqBody CoreKill

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// KillAll is the handler for POST /core0/{id}/core/killall
// Kills all running commands
func (api Core0API) KillAll(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// KVMList is the handler for GET /core0/{id}/kvm
// List kvmdomain
func (api Core0API) KVMList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// KVMCreate is the handler for POST /core0/{id}/kvm
// Create a new kvmdomain
func (api Core0API) KVMCreate(w http.ResponseWriter, r *http.Request) {
	var reqBody KVMCreate

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// KVMGet is the handler for GET /core0/{id}/kvm/{domainid}
// Get detailed domain object
func (api Core0API) KVMGet(w http.ResponseWriter, r *http.Request) {
	var respBody KVMDomain
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// KVMDelete is the handler for DELETE /core0/{id}/kvm/{domainid}
// Delete Domain
func (api Core0API) KVMDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// KVMStart is the handler for POST /core0/{id}/kvm/{domainid}/start
// Start the kvmdomain
func (api Core0API) KVMStart(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// KVMStop is the handler for POST /core0/{id}/kvm/{domainid}/stop
// Gracefully stop the kvmdomain
func (api Core0API) KVMStop(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// KVMDestroy is the handler for POST /core0/{id}/kvm/{domainid}/destroy
// Destroy the kvmdomain
func (api Core0API) KVMDestroy(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// KVMPause is the handler for POST /core0/{id}/kvm/{domainid}/pause
// Pause the kvmdomain
func (api Core0API) KVMPause(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CPUInfo is the handler for GET /core0/{id}/info/cpu
func (api Core0API) CPUInfo(w http.ResponseWriter, r *http.Request) {
	cl := GetConnection(r)
	info := client.Info(cl)
	result, err := info.CPU()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var respBody []CPUInfo

	for i, cpu := range result {
		var info CPUInfo
		info.CacheSize = int(cpu.CacheSize)
		info.CPUInfo = i
		info.CoreId = cpu.CoreID
		info.Cores = int(cpu.Cores)
		info.Family = cpu.Family
		info.Flags = cpu.Flags
		info.Mhz = cpu.Mhz

		respBody = append(respBody, info)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&respBody)
}

// DiskInfo is the handler for GET /core0/{id}/info/disk
func (api Core0API) DiskInfo(w http.ResponseWriter, r *http.Request) {
	cl := GetConnection(r)
	info := client.Info(cl)
	result, err := info.Disk()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var respBody []DiskInfo
	for _, disk := range result {
		var info DiskInfo
		info.Device = disk.Device
		info.Fstype = disk.Fstype
		info.Mountpoint = disk.Mountpoint
		info.Opts = disk.Opts
		respBody = append(respBody, info)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&respBody)
}

// MemInfo is the handler for GET /core0/{id}/info/mem
func (api Core0API) MemInfo(w http.ResponseWriter, r *http.Request) {
	cl := GetConnection(r)
	info := client.Info(cl)
	result, err := info.Mem()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var respBody MemInfo
	respBody.Active = int(result.Active)
	respBody.Available = int(result.Available)
	respBody.Buffers = int(result.Buffers)
	respBody.Cached = int(result.Cached)
	respBody.Free = int(result.Free)
	respBody.Inactive = int(result.Inactive)
	respBody.Total = int(result.Total)
	respBody.Used = int(result.Used)
	respBody.UsedPercent = result.UsedPercent
	respBody.Wired = int(result.Wired)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&respBody)
}

// NicInfo is the handler for GET /core0/{id}/info/nic
func (api Core0API) NicInfo(w http.ResponseWriter, r *http.Request) {
	cl := GetConnection(r)
	info := client.Info(cl)
	result, err := info.Nic()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var respBody []NicInfo
	for _, nic := range result {
		var info NicInfo
		for _, addr := range nic.Addrs {
			info.Addrs = append(info.Addrs, addr.Addr)
		}
		// info.Addrs = nic.Addrs
		info.Flags = nic.Flags
		info.Hardwareaddr = nic.HardwareAddr
		info.Mtu = nic.MTU
		info.Name = nic.Name
		respBody = append(respBody, info)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&respBody)
}

// OSInfo is the handler for GET /core0/{id}/info/os
func (api Core0API) OSInfo(w http.ResponseWriter, r *http.Request) {
	cl := GetConnection(r)
	info := client.Info(cl)
	result, err := info.OS()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var respBody OSInfo
	respBody.BootTime = int(result.BootTime)
	respBody.Hostname = result.Hostname
	respBody.Os = result.OS
	respBody.Platform = result.Platform
	respBody.PlatformFamily = result.PlatformFamily
	respBody.PlatformVersion = result.PlatformVersion
	respBody.Procs = int(result.Procs)
	respBody.Uptime = int(result.Uptime)
	respBody.VirtualizationRole = result.VirtualizationRole
	respBody.VirtualizationSystem = result.VirtualizationSystem

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&respBody)
}

// ProcessList is the handler for GET /core0/{id}/process
// Get Processes
func (api Core0API) ProcessList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	cl := GetConnection(r)
	core := client.Core(cl)

	processes, err := core.Processes()
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	for _, ps := range processes {
		respBody = append(respBody,
			Location{
				Name: ps.Command.Command,
				Url:  Url(r, "process", ps.Command.ID),
				Id:   ps.Command.ID,
			},
		)
	}

	json.NewEncoder(w).Encode(&respBody)
	// w.Header().Set("key","value")
}

// ProcessGet is the handler for GET /core0/{id}/process/{proccessid}
// Get process details
func (api Core0API) ProcessGet(w http.ResponseWriter, r *http.Request) {
	var respBody Process
	json.NewEncoder(w).Encode(&respBody)
	cl := GetConnection(r)
	core := client.Core(cl)

	vars := mux.Vars(r)
	ps, err := core.Process(client.Job(vars["processid"]))
	if err != nil {
		WriteError(w, http.StatusNotFound, err)
		return
	}

	respBody.Cpu = ps.CPU
	respBody.Rss = float64(ps.RSS)
	respBody.Swap = float64(ps.Swap)
	respBody.Vms = float64(ps.VMS)

	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// ProcessKill is the handler for DELETE /core0/{id}/process/{proccessid}
// Kill Process
func (api Core0API) ProcessKill(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BridgeList is the handler for GET /core0/{id}/bridge
// List bridges
func (api Core0API) BridgeList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BridgeCreate is the handler for POST /core0/{id}/bridge
// Creates a new bridge
func (api Core0API) BridgeCreate(w http.ResponseWriter, r *http.Request) {
	var reqBody BridgeCreate

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BridgeGet is the handler for GET /core0/{id}/bridge/{bridgeid}
// Get bridge details
func (api Core0API) BridgeGet(w http.ResponseWriter, r *http.Request) {
	var respBody Bridge
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BridgeDelete is the handler for DELETE /core0/{id}/bridge/{bridgeid}
// Remove bridge
func (api Core0API) BridgeDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// DiskList is the handler for GET /core0/{id}/disk
// List blockdevices present in the Core0
func (api Core0API) DiskList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// DiskGet is the handler for GET /core0/{id}/disk/{devicenameOrdiskserial}
// Get disk details
func (api Core0API) DiskGet(w http.ResponseWriter, r *http.Request) {
	var respBody DiskExtInfo
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// DiskMakeTable is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/mktable
// Creates a partition table on a blockdevice
func (api Core0API) DiskMakeTable(w http.ResponseWriter, r *http.Request) {
	var reqBody DiskMKTable

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// DiskMount is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/mount
// Mounts the disk
func (api Core0API) DiskMount(w http.ResponseWriter, r *http.Request) {
	var reqBody DiskMount

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// DiskUmount is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/umount
// Unmount the disk
func (api Core0API) DiskUmount(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// DiskPartitionsList is the handler for GET /core0/{id}/disk/{devicenameOrdiskserial}/partitions
// Lists partitions
func (api Core0API) DiskPartitionsList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// DiskPartitionCreate is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/partitions
// Create a new partition
func (api Core0API) DiskPartitionCreate(w http.ResponseWriter, r *http.Request) {
	var reqBody DiskCreatePartition

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// DiskPartitionGet is the handler for GET /core0/{id}/disk/{devicenameOrdiskserial}/partitions/{partitionid}
// Gets partition info. (TODO Needs further speccing)
func (api Core0API) DiskPartitionGet(w http.ResponseWriter, r *http.Request) {
	var respBody string
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// DiskPartitionDelete is the handler for DELETE /core0/{id}/disk/{devicenameOrdiskserial}/partitions/{partitionid}
// Removes a partition
func (api Core0API) DiskPartitionDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// DiskPartitionUmount is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/partitions/{partitionid}/umount
// Unmount the partition
func (api Core0API) DiskPartitionUmount(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// DiskPartitionMount is the handler for POST /core0/{id}/disk/{devicenameOrdiskserial}/partitions/{partitionid}/mount
// Mounts the partition
func (api Core0API) DiskPartitionMount(w http.ResponseWriter, r *http.Request) {
	var reqBody DiskMount

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// FileDownload is the handler for GET /core0/{id}/filesystem
// Download file from Core0
func (api Core0API) FileDownload(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// FileUpload is the handler for POST /core0/{id}/filesystem
// Upload file to Core0
func (api Core0API) FileUpload(w http.ResponseWriter, r *http.Request) {
	var reqBody WriteFile

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// FileDelete is the handler for DELETE /core0/{id}/filesystem
// Delete file from Core0
func (api Core0API) FileDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody DeleteFile

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// ZerotierList is the handler for GET /core0/{id}/zerotier
// List running zerotier networks
func (api Core0API) ZerotierList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// ZerotierGet is the handler for GET /core0/{id}/zerotier/{zerotierid}
// Get zerotier network details
func (api Core0API) ZerotierGet(w http.ResponseWriter, r *http.Request) {
	var respBody ZeroTier
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// ZerotierJoin is the handler for POST /core0/{id}/zerotier/{zerotierid}
// Join zerotier network
func (api Core0API) ZerotierJoin(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// ZerotierUnjoin is the handler for DELETE /core0/{id}/zerotier/{zerotierid}
// Delete zerotier network
func (api Core0API) ZerotierUnjoin(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXFileDownload is the handler for GET /core0/{id}/coreX/{coreXid}/filesystem
// Download file from Core0
func (api Core0API) CXFileDownload(w http.ResponseWriter, r *http.Request) {
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXFileUpload is the handler for POST /core0/{id}/coreX/{coreXid}/filesystem
// Upload file to Core0
func (api Core0API) CXFileUpload(w http.ResponseWriter, r *http.Request) {
	var reqBody WriteFile

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXFileDelete is the handler for DELETE /core0/{id}/coreX/{coreXid}/filesystem
// Delete file from Core0
func (api Core0API) CXFileDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody DeleteFile

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXKill is the handler for POST /core0/{id}/coreX/{coreXid}/core/kill
// Kill a process / command
func (api Core0API) CXKill(w http.ResponseWriter, r *http.Request) {
	var reqBody CoreKill

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXKillAll is the handler for POST /core0/{id}/coreX/{coreXid}/core/killall
// Kills all running commands
func (api Core0API) CXKillAll(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXStateGet is the handler for GET /core0/{id}/coreX/{coreXid}/core/state
// The aggregated consumption of coreX + all processes (cpu, memory, etc...)
func (api Core0API) CXStateGet(w http.ResponseWriter, r *http.Request) {
	var respBody CoreStateResult
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXPing is the handler for POST /core0/{id}/coreX/{coreXid}/core/ping
// Execute a ping command to this CoreX
func (api Core0API) CXPing(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXSystem is the handler for POST /core0/{id}/coreX/{coreXid}/core/system
// Execute a new process  on this CoreX
func (api Core0API) CXSystem(w http.ResponseWriter, r *http.Request) {
	var reqBody CoreSystem

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXCPUInfo is the handler for GET /core0/{id}/coreX/{coreXid}/info/cpu
func (api Core0API) CXCPUInfo(w http.ResponseWriter, r *http.Request) {
	var respBody []CPUInfo
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXMemInfo is the handler for GET /core0/{id}/coreX/{coreXid}/info/mem
func (api Core0API) CXMemInfo(w http.ResponseWriter, r *http.Request) {
	var respBody MemInfo
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXNicInfo is the handler for GET /core0/{id}/coreX/{coreXid}/info/nic
func (api Core0API) CXNicInfo(w http.ResponseWriter, r *http.Request) {
	var respBody []NicInfo
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXOSInfo is the handler for GET /core0/{id}/coreX/{coreXid}/info/os
func (api Core0API) CXOSInfo(w http.ResponseWriter, r *http.Request) {
	var respBody OSInfo
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXProcessList is the handler for GET /core0/{id}/coreX/{coreXid}/process
// Get Processes
func (api Core0API) CXProcessList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXProcessGet is the handler for GET /core0/{id}/coreX/{coreXid}/process/{proccessid}
// Get process details
func (api Core0API) CXProcessGet(w http.ResponseWriter, r *http.Request) {
	var respBody Process
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXProcessKill is the handler for DELETE /core0/{id}/coreX/{coreXid}/process/{proccessid}
// Kill Process
func (api Core0API) CXProcessKill(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BtrfsList is the handler for GET /core0/{id}/btrfs
// List btrfs filesystems
func (api Core0API) BtrfsList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BtrfsCreate is the handler for POST /core0/{id}/btrfs
// Create a new btrfs filesystem
func (api Core0API) BtrfsCreate(w http.ResponseWriter, r *http.Request) {
	var reqBody BtrfsCreate

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BtrfsGet is the handler for GET /core0/{id}/btrfs/{btrfsUuid}
// Get btrfs filesystem details
func (api Core0API) BtrfsGet(w http.ResponseWriter, r *http.Request) {
	var respBody Btrfs
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BtrfsAddDevice is the handler for POST /core0/{id}/btrfs/{btrfsUuid}
// Add a device to an existing btrfs filesystem
func (api Core0API) BtrfsAddDevice(w http.ResponseWriter, r *http.Request) {
	var reqBody BtrfsAddDevice

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BtrfsSubvolList is the handler for GET /core0/{id}/btrfs/{btrfsUuid}/subvolumes
// Btrfs List Subvolumes
func (api Core0API) BtrfsSubvolList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BtrfsSubvolGet is the handler for GET /core0/{id}/btrfs/{btrfsUuid}/subvolumes/{subvolumepath}
// Gets subvolume details
func (api Core0API) BtrfsSubvolGet(w http.ResponseWriter, r *http.Request) {
	var respBody BtrfsSubVolume
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BtrfsSubvolCreate is the handler for POST /core0/{id}/btrfs/{btrfsUuid}/subvolumes/{subvolumepath}
// Create a new subvolume
func (api Core0API) BtrfsSubvolCreate(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// BtrfsSubvolDelete is the handler for DELETE /core0/{id}/btrfs/{btrfsUuid}/subvolumes/{subvolumepath}
// Delete subvolume
func (api Core0API) BtrfsSubvolDelete(w http.ResponseWriter, r *http.Request) {
	var reqBody Command

	// decode request
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(400)
		return
	}

	// validate request
	if err := reqBody.Validate(); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}
	var respBody Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXCommandsList is the handler for GET /core0/{id}/coreX/{coreXid}/command
// List running commands
func (api Core0API) CXCommandsList(w http.ResponseWriter, r *http.Request) {
	var respBody []Location
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}

// CXCommandGet is the handler for GET /core0/{id}/coreX/{coreXid}/command/{commandid}
func (api Core0API) CXCommandGet(w http.ResponseWriter, r *http.Request) {
	var respBody CommandResult
	json.NewEncoder(w).Encode(&respBody)
	// uncomment below line to add header
	// w.Header().Set("key","value")
}
