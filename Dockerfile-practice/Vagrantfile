# This block starts the Vagrant configuration and
# specifies the configuration version. "2" refers to Vagrant 1.1 and above.
Vagrant.configure("2") do |config|

    # Define the first VM
    config.vm.define "web" do |web|
      web.vm.box = "ubuntu/bionic64"
      
      # Network settings
      web.vm.network "private_network", ip: "192.168.33.10"
      web.vm.network "forwarded_port", guest: 80, host: 8080
  
      # Synced folder
      web.vm.synced_folder "../data", "/vagrant_data"
  
      # Provider settings
      web.vm.provider "virtualbox" do |vb|
        vb.name = "web-server"
        vb.memory = "2048"
        vb.cpus = 2
      end
  
      # Provision with a shell script
      web.vm.provision "shell", inline: <<-SHELL
        sudo apt-get update
        sudo apt-get install -y apache2
      SHELL
  
      # Provision with Ansible
      web.vm.provision "ansible" do |ansible|
        ansible.playbook = "playbook.yml"
      end
    end
  
    # Define the second VM
    config.vm.define "db" do |db|
      db.vm.box = "ubuntu/bionic64"
      
      # Network settings
      db.vm.network "private_network", ip: "192.168.33.11"
  
      # Provider settings
      db.vm.provider "virtualbox" do |vb|
        vb.name = "db-server"
        vb.memory = "1024"
        vb.cpus = 1
      end
  
      # Provision with a shell script
      db.vm.provision "shell", inline: <<-SHELL
        sudo apt-get update
        sudo apt-get install -y mysql-server
        sudo mysql_secure_installation
      SHELL
    end
  
    # Global settings
    config.vm.provider "virtualbox" do |vb|
      vb.customize ["modifyvm", :id, "--ioapic", "on"]
    end
  
    config.vm.provision "shell", inline: <<-SHELL
      echo "This is a global provisioner that runs on all VMs."
    SHELL
  
  end
  

#The following code is for a LAMP stack. 
#LAMP is a popular software stack used for web development.
#The acronym stands for Linux, Apache, MySQL (or MariaDB), and PHP (or Perl/Python).


# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
    # Use Ubuntu 20.04 as the base box
    config.vm.box = "ubuntu/focal64"
  
    # Assign a static IP to the VM (optional)
    config.vm.network "private_network", ip: "192.168.56.10"
  
    # Configure the VM to use 2 GB of RAM and 2 CPUs (optional)
    config.vm.provider "virtualbox" do |vb|
      vb.memory = "2048"
      vb.cpus = 2
    end
  
    # Provision the VM with a shell script
    config.vm.provision "shell", inline: <<-SHELL
      # Update package index
      sudo apt-get update
  
      # Install Apache
      sudo apt-get install -y apache2
  
      # Install MySQL
      sudo apt-get install -y mariadb-server mariadb-client
  
      # Secure MySQL installation
      sudo mysql -e "UPDATE mysql.user SET Password = PASSWORD('root') WHERE User = 'root'"
      sudo mysql -e "DELETE FROM mysql.user WHERE User=''"
      sudo mysql -e "DELETE FROM mysql.user WHERE User='root' AND Host NOT IN ('localhost', '127.0.0.1', '::1')"
      sudo mysql -e "DROP DATABASE IF EXISTS test"
      sudo mysql -e "DELETE FROM mysql.db WHERE Db='test' OR Db='test\\_%'"
      sudo mysql -e "FLUSH PRIVILEGES"
  
      # Install PHP
      sudo apt-get install -y php libapache2-mod-php php-mysql
  
      # Configure Apache to prefer PHP files
      sudo sed -i 's/index.html/index.php index.html/g' /etc/apache2/mods-enabled/dir.conf
  
      # Restart Apache to apply changes
      sudo systemctl restart apache2
  
      # Create a PHP info page for testing
      echo "<?php phpinfo(); ?>" | sudo tee /var/www/html/info.php
  
      # Install phpMyAdmin (optional)
      echo "phpmyadmin phpmyadmin/dbconfig-install boolean true" | sudo debconf-set-selections
      echo "phpmyadmin phpmyadmin/app-password-confirm password root" | sudo debconf-set-selections
      echo "phpmyadmin phpmyadmin/mysql/admin-pass password root" | sudo debconf-set-selections
      echo "phpmyadmin phpmyadmin/mysql/app-pass password root" | sudo debconf-set-selections
      echo "phpmyadmin phpmyadmin/reconfigure-webserver multiselect apache2" | sudo debconf-set-selections
      sudo apt-get install -y phpmyadmin
    SHELL
  end
  