ports = node['robi']['api']['ports']
image = node['robi']['api']['image']
script "start_api_container" do
	interpreter "bash"
    user "root"
	code <<-EOH
		docker run -d #{ports} --name=api #{image}
	EOH
end
