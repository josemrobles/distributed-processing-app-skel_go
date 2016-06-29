script "pull_api_image" do
	interpreter "bash"
	user "root"
	code <<-EOH
		docker pull #{node['robi']['api']['image']}
	EOH
end
