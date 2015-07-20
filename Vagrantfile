# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
    config.vm.box = "hashicorp/precise64"
    config.vm.network :private_network, ip: "192.168.33.20"
    config.vm.network "forwarded_port", guest: 8080, host: 3000
    config.vm.synced_folder "./", "/home/vagrant/go/src/godemo"
    config.vm.provision :shell, :inline => <<-EOF
    ZONE=`date +%Z`
    if [ "$ZONE" != "JST" ] ; then
        echo "Asia/Tokyo" > /etc/timezone
        dpkg-reconfigure -f noninteractive tzdata
        fi
    EOF
    config.vm.provision "ansible" do |ansible|
        ansible.playbook = "provisioning/playbook.yml"
        ansible.inventory_path = "provisioning/hosts"
    end
end
