include_recipe 'apt'

service "docker" do
	action :nothing
	supports :start => true, :stop => true, :restart => true
end

execute "apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D"

directory "/etc/apt/sources.list.d"

execute "echo deb https://apt.dockerproject.org/repo $(lsb_release -si| tr '[:upper:]' '[:lower:]' )-$(lsb_release --codename | cut -f2) main  > /etc/apt/sources.list.d/docker.list"

execute 'apt-get update' do
	command 'apt-get update'
	ignore_failure true
end

package 'docker-engine' do
	action :install
end
