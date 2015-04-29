# -*- mode: ruby -*-
# vi: set ft=ruby :

# ssh-agent must be aware of your identity or vagrant ssh forwarding will fail
def sshCheck
  output=`ssh-add -L 2>&1`
  if 0 != $?
    if output =~ /Could not open a connection/
      puts 'error: start your ssh-agent'
      puts 'bash: eval `ssh-agent -s`'
      puts 'rc:   eval `{ssh-agent -s}'
    elsif output =~ /The agent has no identities/
      puts 'error: add your ssh identity to the agent'
      puts 'ssh-add [key-path]'
    else
      puts 'error: something is wrong with your ssh-agent'
    end

    return false
  end

  return true
end

exit 1 if false == sshCheck

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "golang-lab"
  config.vm.box_url = "http://cloud-images.ubuntu.com/vagrant/trusty/current/trusty-server-cloudimg-amd64-vagrant-disk1.box"
  config.vm.network "private_network", ip: "192.168.13.50"
  config.ssh.forward_agent = true
  config.vm.synced_folder ".", "/opt/golang-lab"
  config.vm.provider "virtualbox" do |v|
      v.memory = 2048
      v.cpus = 2
  end
  config.vm.provision "shell", path: "./vagrant/provision.sh"
end
