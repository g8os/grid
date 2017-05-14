from JumpScale import j


def init(job):
    service = job.service
    node = j.sal.g8os.get_node(
        addr=service.model.data.redisAddr,
        port=service.model.data.redisPort,
        password=service.model.data.redisPassword or None,
    )

    job.logger.info("create storage pool for fuse cache")
    poolname = "{}_fscache".format(service.name)

    storagepool = node.ensure_persistance(poolname)
    storagepool.ays.create(service.aysrepo)


def getAddresses(job):
    service = job.service
    networks = service.producers.get('network', [])
    networkmap = {}
    for network in networks:
        job = network.getJob('getAddresses', args={'node_name': service.name})
        networkmap[network.name] = j.tools.async.wrappers.sync(job.execute())
    return networkmap


def install(job):
    # at each boot recreate the complete state in the system
    service = job.service
    node = j.sal.g8os.get_node(
        addr=service.model.data.redisAddr,
        port=service.model.data.redisPort,
        password=service.model.data.redisPassword or None,
    )

    job.logger.info("mount storage pool for fuse cache")
    poolname = "{}_fscache".format(service.name)
    node.ensure_persistance(poolname)

    job.logger.info("configure networks")
    for network in service.producers.get('network', []):
        job = network.getJob('configure', args={'node_name': service.name})
        j.tools.async.wrappers.sync(job.execute())


def monitor(job):
    import redis
    service = job.service
    addr = service.model.data.redisAddr
    node = j.clients.g8core.get(addr, testConnectionAttempts=0)
    node.timeout = 15
    try:
        state = node.ping()
    except redis.ConnectionError as e:
        state = False

    if state:
        service.model.data.status = 'running'
    else:
        service.model.data.status = 'halted'
    service.saveAll()


def reboot(job):
    service = job.service
    job.logger.info("reboot node {}".format(service))
    node = j.sal.g8os.get_node(
        addr=service.model.data.redisAddr,
        port=service.model.data.redisPort,
        password=service.model.data.redisPassword or None,
    )
    node.client.raw('core.reboot', {})


def uninstall(job):
    service = job.service
    bootstraps = service.aysrepo.servicesFind(actor='bootstrap.g8os')
    if bootstraps:
        j.tools.async.wrappers.sync(bootstraps[0].getJob('delete_node', args={'node_name': service.name}).execute())
