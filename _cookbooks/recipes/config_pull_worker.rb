script "pull_worker_image" do
	interpreter "bash"
	user "root"
	code <<-EOH
		docker pull #{node['robi']['worker']['image']}
	EOH
end
