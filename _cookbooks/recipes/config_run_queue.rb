ports = node['robi']['queue']['ports']
image = node['robi']['queue']['image']
script "start_queue_container" do
	interpreter "bash"
    user "root"
	code <<-EOH
		docker run -d #{ports} --name=queue #{image}
	EOH
end
