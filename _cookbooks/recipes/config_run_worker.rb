ports = node['robi']['worker']['ports']
image = node['robi']['worker']['image']
script "start_worker_container" do
	interpreter "bash"
    user "root"
	code <<-EOH
		docker run -d #{ports} --name=worker #{image}
	EOH
end
