script "pull_queue_image" do
	interpreter "bash"
	user "root"
	code <<-EOH
		docker pull #{node['robi']['queue']['image']}
	EOH
end
