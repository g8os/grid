from js9 import j


def install(job):
    import yaml
    from io import BytesIO

    service = job.service
    container = get_container(service, job.context['token'])

    configpath = "/etc/etcd_{}.config".format(service.name)

    config = {
        "name": service.name,
        "initial-advertise-peer-urls": "http://{}".format(service.model.data.serverBind),
        "listen-peer-urls": "http://{}".format(service.model.data.serverBind),
        "listen-client-urls": "http://{}".format(service.model.data.clientBind),
        "advertise-client-urls": "http://{}".format(service.model.data.clientBind),
        "initial-cluster": ",".join(service.model.data.peers),
        "initial-cluster-state": "new"
    }
    yamlconfig = yaml.safe_dump(config, default_flow_style=False)
    configstream = BytesIO(yamlconfig.encode('utf8'))
    configstream.seek(0)
    container.client.filesystem.upload(configpath, configstream)

    cmd = '/bin/etcd --config-file %s' % configpath
    container.client.system(cmd, id="{}.{}".format(service.model.role, service.name))

    if not container.is_port_listening(int(service.model.data.serverBind.split(":")[1])):
        raise j.exceptions.RuntimeError('Failed to start etcd server: {}'.format(service.name))

    service.model.data.status = "running"


def start(job):
    service = job.service
    j.tools.async.wrappers.sync(service.executeAction('install', context=job.context))


def get_container(service, password):
    from zeroos.orchestrator.sal.Container import Container
    return Container.from_ays(service.parent, password)


def watchdog_handler(job):
    import asyncio
    service = job.service
    loop = j.atyourservice.server.loop
    if service.model.data.status == 'running':
        asyncio.ensure_future(job.service.executeAction('start', context=job.context), loop=loop)
