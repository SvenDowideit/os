package integration

import . "github.com/cpuguy83/check"

func (s *QemuSuite) TestCustomDocker(c *C) {
	c.Parallel()
	err := s.RunQemu("--cloud-config", "./tests/assets/test_05/cloud-config.yml")
	c.Assert(err, IsNil)

	s.CheckCall(c, `
set -ex

docker version | grep 1.10.3

sudo ros engine list | grep 1.10.3 | grep current
docker run -d --restart=always nginx
docker ps | grep nginx`)

	s.CheckCall(c, `
set -ex

sudo ros engine switch docker-1.11.2
/usr/sbin/wait-for-docker
docker version | grep 1.11.2
sudo ros engine list | grep 1.11.2 | grep current
docker ps | grep nginx`)

	s.Reboot(c)

	s.CheckCall(c, `
set -ex

docker version | grep 1.11.2
sudo ros engine list | grep 1.11.2 | grep current
docker ps | grep nginx`)
}

func (s *QemuSuite) TestCustomDockerInPersistentConsole(c *C) {
	c.Parallel()
	err := s.RunQemu("--cloud-config", "./tests/assets/test_25/cloud-config.yml")
	c.Assert(err, IsNil)

	s.CheckCall(c, `
set -ex

apt-get --version
docker version | grep 1.10.3
sudo ros engine list | grep 1.10.3 | grep current
docker run -d --restart=always nginx
docker ps | grep nginx`)

	s.CheckCall(c, `
set -ex

sudo ros engine switch docker-1.11.2
/usr/sbin/wait-for-docker
docker version | grep 1.11.2
sudo ros engine list | grep 1.11.2 | grep current
docker ps | grep nginx`)

	s.Reboot(c)

	s.CheckCall(c, `
set -ex

docker version | grep 1.11.2
sudo ros engine list | grep 1.11.2 | grep current
docker ps | grep nginx`)
}
