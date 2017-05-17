package node

//This file is auto-generated by go-raml
//Do not edit this file by hand since it will be overwritten during the next generation

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NodeInterface is interface for /node root endpoint
type NodesInterface interface { // ListNodes is the handler for GET /node
	// List Nodes
	ListNodes(http.ResponseWriter, *http.Request)
	// GetNode is the handler for GET /node/{nodeid}
	// Get detailed information of a node
	GetNode(http.ResponseWriter, *http.Request)
	// DeleteNode is the handler for DELETE /node/{nodeid}
	// Delete a node
	DeleteNode(http.ResponseWriter, *http.Request)
	// GetStoragePools is the handler for GET /node/{nodeid}/storagepool
	// List storage pools present in the node
	ListStoragePools(http.ResponseWriter, *http.Request)
	// CreateStoragePool is the handler for POST /node/{nodeid}/storagepool
	// Create a new storage pool
	CreateStoragePool(http.ResponseWriter, *http.Request)
	// GetStoragePoolInfo is the handler for GET /node/{nodeid}/storagepool/{storagepoolname}
	// Get detailed information of this storage pool
	GetStoragePoolInfo(http.ResponseWriter, *http.Request)
	// DeleteStoragePool is the handler for DELETE /node/{nodeid}/storagepool/{storagepoolname}
	// Delete the storage pool
	DeleteStoragePool(http.ResponseWriter, *http.Request)
	// ListStoragePoolDevices is the handler for GET /node/{nodeid}/storagepool/{storagepoolname}/devices
	// Lists the devices in the storage pool
	ListStoragePoolDevices(http.ResponseWriter, *http.Request)
	// CreateStoragePoolDevices is the handler for POST /node/{nodeid}/storagepool/{storagepoolname}/devices
	// Add extra devices to this storage pool
	CreateStoragePoolDevices(http.ResponseWriter, *http.Request)
	// GetStoragePoolDeviceInfo is the handler for GET /node/{nodeid}/storagepool/{storagepoolname}/devices/{deviceuuid}
	// Get information of the device
	GetStoragePoolDeviceInfo(http.ResponseWriter, *http.Request)
	// DeleteStoragePoolDevice is the handler for DELETE /node/{nodeid}/storagepool/{storagepoolname}/devices/{deviceuuid}
	// Removes the device from the storagepool
	DeleteStoragePoolDevice(http.ResponseWriter, *http.Request)
	// ListFilesystems is the handler for GET /node/{nodeid}/storagepool/{storagepoolname}/filesystems
	// List filesystems
	ListFilesystems(http.ResponseWriter, *http.Request)
	// CreateFilesystem is the handler for POST /node/{nodeid}/storagepool/{storagepoolname}/filesystems
	// Create a new filesystem
	CreateFilesystem(http.ResponseWriter, *http.Request)
	// GetFilesystemInfo is the handler for GET /node/{nodeid}/storagepool/{storagepoolname}/filesystems/{filesystemname}
	// Get detailed filesystem information
	GetFilesystemInfo(http.ResponseWriter, *http.Request)
	// DeleteFilesystem is the handler for DELETE /node/{nodeid}/storagepool/{storagepoolname}/filesystems/{filesystemname}
	// Delete filesystem
	DeleteFilesystem(http.ResponseWriter, *http.Request)
	// ListFilesystemSnapshots is the handler for GET /node/{nodeid}/storagepool/{storagepoolname}/filesystems/{filesystemname}/snapshot
	// List snapshots of this filesystem
	ListFilesystemSnapshots(http.ResponseWriter, *http.Request)
	// CreateSnapshot is the handler for POST /node/{nodeid}/storagepool/{storagepoolname}/filesystems/{filesystemname}/snapshots
	// Create a new readonly filesystem of the current state of the vdisk
	CreateSnapshot(http.ResponseWriter, *http.Request)
	// GetFilesystemSnapshotInfo is the handler for GET /node/{nodeid}/storagepool/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}
	// Get detailed information on the snapshot
	GetFilesystemSnapshotInfo(http.ResponseWriter, *http.Request)
	// DeleteFilesystemSnapshot is the handler for DELETE /node/{nodeid}/storagepool/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}
	// Delete snapshot
	DeleteFilesystemSnapshot(http.ResponseWriter, *http.Request)
	// RollbackFilesystemSnapshot is the handler for POST /node/{nodeid}/storagepool/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}/rollback
	// Rollback the filesystem to the state at the moment the snapshot was taken
	RollbackFilesystemSnapshot(http.ResponseWriter, *http.Request)
	// ListVMs is the handler for GET /node/{nodeid}/vms
	// List VMs
	ListVMs(http.ResponseWriter, *http.Request)
	// CreateVM is the handler for POST /node/{nodeid}/vms
	// Creates the VM
	CreateVM(http.ResponseWriter, *http.Request)
	// GetVM is the handler for GET /node/{nodeid}/vms/{vmid}
	// Get detailed virtual machine object
	GetVM(http.ResponseWriter, *http.Request)
	// UpdateVM is the handler for PUT /node/{nodeid}/vms/{vmid}
	// Updates virtual machine object
	UpdateVM(http.ResponseWriter, *http.Request)
	// DeleteVM is the handler for DELETE /node/{nodeid}/vms/{vmid}
	// Deletes the VM
	DeleteVM(http.ResponseWriter, *http.Request)
	// MigrateVM is the handler for POST /node/{nodeid}/vms/{vmid}/migrate
	// Migrate the VM to another host
	MigrateVM(http.ResponseWriter, *http.Request)
	// GetVMInfo is the handler for GET /node/{nodeid}/vms/{vmid}/info
	// Get statistical information about the virtual machine.
	GetVMInfo(http.ResponseWriter, *http.Request)
	// StartVM is the handler for POST /node/{nodeid}/vms/{vmid}/start
	// Starts the VM
	StartVM(http.ResponseWriter, *http.Request)
	// StopVM is the handler for POST /node/{nodeid}/vms/{vmid}/stop
	// Stops the VM
	StopVM(http.ResponseWriter, *http.Request)
	// PauseVM is the handler for POST /node/{nodeid}/vms/{vmid}/pause
	// Pauses the VM
	PauseVM(http.ResponseWriter, *http.Request)
	// ResumeVM is the handler for POST /node/{nodeid}/vms/{vmid}/resume
	// Resumes the VM
	ResumeVM(http.ResponseWriter, *http.Request)
	// ShutdownVM is the handler for POST /node/{nodeid}/vms/{vmid}/shutdown
	// Gracefully shutdown the VM
	ShutdownVM(http.ResponseWriter, *http.Request)
	// ListNodeJobs is the handler for GET /node/{nodeid}/jobs
	// List running jobs
	ListNodeJobs(http.ResponseWriter, *http.Request)
	// KillAllNodeJobs is the handler for DELETE /node/{nodeid}/jobs
	// Kills all running jobs
	KillAllNodeJobs(http.ResponseWriter, *http.Request)
	// GetNodeJob is the handler for GET /node/{nodeid}/jobs/{jobid}
	// Get the details of a submitted job
	GetNodeJob(http.ResponseWriter, *http.Request)
	// KillNodeJob is the handler for DELETE /node/{nodeid}/jobs/{jobid}
	// Kills the job
	KillNodeJob(http.ResponseWriter, *http.Request)
	// GetMemInfo is the handler for GET /node/{nodeid}/mem
	// Get detailed information about the memory in the node
	GetMemInfo(http.ResponseWriter, *http.Request)
	// GetNicInfo is the handler for GET /node/{nodeid}/nics
	// Get detailed information about the network interfaces in the node
	GetNicInfo(http.ResponseWriter, *http.Request)
	// ListBridges is the handler for GET /node/{nodeid}/bridges
	// List bridges
	ListBridges(http.ResponseWriter, *http.Request)
	// CreateBridge is the handler for POST /node/{nodeid}/bridges
	// Creates a new bridge
	CreateBridge(http.ResponseWriter, *http.Request)
	// GetBridge is the handler for GET /node/{nodeid}/bridges/{bridgeid}
	// Get bridge details
	GetBridge(http.ResponseWriter, *http.Request)
	// DeleteBridge is the handler for DELETE /node/{nodeid}/bridges/{bridgeid}
	// Remove bridge
	DeleteBridge(http.ResponseWriter, *http.Request)
	// GetNodeOSInfo is the handler for GET /node/{nodeid}/info
	// Get detailed information of the os of the node
	GetNodeOSInfo(http.ResponseWriter, *http.Request)
	// ListZerotier is the handler for GET /node/{nodeid}/zerotiers
	// List running Zerotier networks
	ListZerotier(http.ResponseWriter, *http.Request)
	// JoinZerotier is the handler for POST /node/{nodeid}/zerotiers
	// Join Zerotier network
	JoinZerotier(http.ResponseWriter, *http.Request)
	// GetZerotier is the handler for GET /node/{nodeid}/zerotiers/{zerotierid}
	// Get Zerotier network details
	GetZerotier(http.ResponseWriter, *http.Request)
	// ExitZerotier is the handler for DELETE /node/{nodeid}/zerotiers/{zerotierid}
	// Exit the Zerotier network
	ExitZerotier(http.ResponseWriter, *http.Request)
	// ListContainers is the handler for GET /node/{nodeid}/containers
	// List running Containers
	ListContainers(http.ResponseWriter, *http.Request)
	// CreateContainer is the handler for POST /node/{nodeid}/containers
	// Create a new Container
	CreateContainer(http.ResponseWriter, *http.Request)
	// GetContainer is the handler for GET /node/{nodeid}/containers/{containername}
	// Get Container
	GetContainer(http.ResponseWriter, *http.Request)
	// DeleteContainer is the handler for DELETE /node/{nodeid}/containers/{containername}
	// Delete Container instance
	DeleteContainer(http.ResponseWriter, *http.Request)
	// StopContainer is the handler for POST /nodes/{nodeid}/containers/{containername}/stop
	// Stop Container instance
	StopContainer(http.ResponseWriter, *http.Request)
	// StartContainer is the handler for POST /nodes/{nodeid}/containers/{containername}/start
	// Start Container instance
	StartContainer(http.ResponseWriter, *http.Request)
	// ListContainerJobs is the handler for GET /node/{nodeid}/containers/{containername}/jobs
	// List running jobs on the container
	ListContainerJobs(http.ResponseWriter, *http.Request)
	// GetContainerNicInfo is the handler for GET /node/{nodeid}/containers/{containername}/nics
	// List nic info on the container
	GetContainerNicInfo(http.ResponseWriter, *http.Request)
	// GetContainerMemInfo is the handler for GET /node/{nodeid}/containers/{containername}/mems
	// List mem info on the container
	GetContainerMemInfo(http.ResponseWriter, *http.Request)
	// GetContainerCPUInfo is the handler for GET /node/{nodeid}/containers/{containername}/cpus
	// List cpu info on the container
	GetContainerCPUInfo(http.ResponseWriter, *http.Request)
	// SendSignalJob is the handler for POST /node/{nodeid}/containers/{containername}/jobs/{jobid}
	// Send signal to the job
	SendSignalJob(http.ResponseWriter, *http.Request)
	// KillAllContainerJobs is the handler for DELETE /node/{nodeid}/containers/{containername}/jobs
	// Kills all running jobs on the container
	KillAllContainerJobs(http.ResponseWriter, *http.Request)
	// GetContainerJob is the handler for GET /node/{nodeid}/containers/{containername}/jobs/{jobid}
	// Get details of a submitted job on the container
	GetContainerJob(http.ResponseWriter, *http.Request)
	// KillContainerJob is the handler for DELETE /node/{nodeid}/containers/{containername}/jobs/{jobid}
	// Kills the job
	KillContainerJob(http.ResponseWriter, *http.Request)
	// PingContainer is the handler for POST /node/{nodeid}/containers/{containername}/ping
	// Ping this container
	PingContainer(http.ResponseWriter, *http.Request)
	// GetContainerState is the handler for GET /node/{nodeid}/containers/{containername}/state
	// The aggregated consumption of container + all processes (cpu, memory, etc...)
	GetContainerState(http.ResponseWriter, *http.Request)
	// GetContainerOSInfo is the handler for GET /node/{nodeid}/containers/{containername}/info
	// Get detailed information of the container os
	GetContainerOSInfo(http.ResponseWriter, *http.Request)
	// ListContainerProcesses is the handler for GET /node/{nodeid}/containers/{containername}/processes
	// Get running processes in this container
	ListContainerProcesses(http.ResponseWriter, *http.Request)
	// StartContainerProcess is the handler for POST /node/{nodeid}/containers/{containername}/processes
	// Start a new process in this container
	StartContainerProcess(http.ResponseWriter, *http.Request)
	// GetContainerProcess is the handler for GET /node/{nodeid}/containers/{containername}/processes/{processid}
	// Get process details
	GetContainerProcess(http.ResponseWriter, *http.Request)
	// SendSignalProcess is the handler for POST /node/{nodeid}/containers/{containername}/processes/{processid}
	// Send signal to the process
	SendSignalProcess(http.ResponseWriter, *http.Request)
	// KillContainerProcess is the handler for DELETE /node/{nodeid}/containers/{containername}/processes/{processid}
	// Kill Process
	KillContainerProcess(http.ResponseWriter, *http.Request)
	// FileDownload is the handler for GET /node/{nodeid}/containers/{containername}/filesystems
	// Download file from container
	FileDownload(http.ResponseWriter, *http.Request)
	// FileUpload is the handler for POST /node/{nodeid}/containers/{containername}/filesystems
	// Upload file to container
	FileUpload(http.ResponseWriter, *http.Request)
	// FileDelete is the handler for DELETE /node/{nodeid}/containers/{containername}/filesystems
	// Delete file from container
	FileDelete(http.ResponseWriter, *http.Request)
	// ListNodeProcesses is the handler for GET /node/{nodeid}/processes
	// Get Processes
	ListNodeProcesses(http.ResponseWriter, *http.Request)
	// GetNodeProcess is the handler for GET /node/{nodeid}/processes/{processid}
	// Get process details
	GetNodeProcess(http.ResponseWriter, *http.Request)
	// KillNodeProcess is the handler for DELETE /node/{nodeid}/processes/{processid}
	// Kill Process
	KillNodeProcess(http.ResponseWriter, *http.Request)
	// PingNode is the handler for POST /node/{nodeid}/ping
	// Ping this node
	PingNode(http.ResponseWriter, *http.Request)
	// RebootNode is the handler for POST /node/{nodeid}/reboot
	// Immediately reboot the machine.
	RebootNode(http.ResponseWriter, *http.Request)
	// GetCPUInfo is the handler for GET /node/{nodeid}/cpus
	// Get detailed information of all CPUs in the node
	GetCPUInfo(http.ResponseWriter, *http.Request)
	// GetDiskInfo is the handler for GET /node/{nodeid}/disk
	// Get detailed information of all the disks in the node
	GetDiskInfo(http.ResponseWriter, *http.Request)
	// GetNodeState is the handler for GET /node/{nodeid}/state
	// The aggregated consumption of node + all processes (cpu, memory, etc...)
	GetNodeState(http.ResponseWriter, *http.Request)
	// GetNodeMounts is the handler for GET /node/{nodeid}/mounts
	// The mountpoints of the node
	GetNodeMounts(http.ResponseWriter, *http.Request)
}

// NodesInterfaceRoutes is routing for /node root endpoint
func NodesInterfaceRoutes(r *mux.Router, i NodesInterface) {
	r.HandleFunc("/nodes", i.ListNodes).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}", i.GetNode).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}", i.DeleteNode).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/storagepools", i.ListStoragePools).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools", i.CreateStoragePool).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}", i.GetStoragePoolInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}", i.DeleteStoragePool).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/devices", i.ListStoragePoolDevices).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/devices", i.CreateStoragePoolDevices).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/devices/{deviceuuid}", i.GetStoragePoolDeviceInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/devices/{deviceuuid}", i.DeleteStoragePoolDevice).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems", i.ListFilesystems).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems", i.CreateFilesystem).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}", i.GetFilesystemInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}", i.DeleteFilesystem).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots", i.ListFilesystemSnapshots).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots", i.CreateSnapshot).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}", i.GetFilesystemSnapshotInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}", i.DeleteFilesystemSnapshot).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/storagepools/{storagepoolname}/filesystems/{filesystemname}/snapshots/{snapshotname}/rollback", i.RollbackFilesystemSnapshot).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms", i.ListVMs).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/vms", i.CreateVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}", i.GetVM).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}", i.UpdateVM).Methods("PUT")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}", i.DeleteVM).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/migrate", i.MigrateVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/info", i.GetVMInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/start", i.StartVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/stop", i.StopVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/pause", i.PauseVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/resume", i.ResumeVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/vms/{vmid}/shutdown", i.ShutdownVM).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/jobs", i.ListNodeJobs).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/jobs", i.KillAllNodeJobs).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/jobs/{jobid}", i.GetNodeJob).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/jobs/{jobid}", i.KillNodeJob).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/mem", i.GetMemInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/nics", i.GetNicInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/bridges", i.ListBridges).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/bridges", i.CreateBridge).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/bridges/{bridgeid}", i.GetBridge).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/bridges/{bridgeid}", i.DeleteBridge).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/info", i.GetNodeOSInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/zerotiers", i.ListZerotier).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/zerotiers", i.JoinZerotier).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/zerotiers/{zerotierid}", i.GetZerotier).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/zerotiers/{zerotierid}", i.ExitZerotier).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/containers", i.ListContainers).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers", i.CreateContainer).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}", i.GetContainer).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}", i.DeleteContainer).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/start", i.StartContainer).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/stop", i.StopContainer).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/jobs", i.ListContainerJobs).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/jobs", i.KillAllContainerJobs).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/jobs/{jobid}", i.GetContainerJob).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/jobs/{jobid}", i.SendSignalJob).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/jobs/{jobid}", i.KillContainerJob).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/nics", i.GetContainerNicInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/mems", i.GetContainerMemInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/cpus", i.GetContainerCPUInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/ping", i.PingContainer).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/state", i.GetContainerState).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/info", i.GetContainerOSInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/processes", i.ListContainerProcesses).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/processes", i.StartContainerProcess).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/processes/{processid}", i.GetContainerProcess).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/processes/{processid}", i.SendSignalProcess).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/processes/{processid}", i.KillContainerProcess).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/filesystem", i.FileDownload).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/filesystem", i.FileUpload).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/containers/{containername}/filesystem", i.FileDelete).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/processes", i.ListNodeProcesses).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/processes/{processid}", i.GetNodeProcess).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/processes/{processid}", i.KillNodeProcess).Methods("DELETE")
	r.HandleFunc("/nodes/{nodeid}/ping", i.PingNode).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/reboot", i.RebootNode).Methods("POST")
	r.HandleFunc("/nodes/{nodeid}/cpus", i.GetCPUInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/disks", i.GetDiskInfo).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/state", i.GetNodeState).Methods("GET")
	r.HandleFunc("/nodes/{nodeid}/mounts", i.GetNodeMounts).Methods("GET")
}
