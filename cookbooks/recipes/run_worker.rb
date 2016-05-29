node[:apps].each do |name, image|  
  script "run_app_#{name}_container" do
    interpreter "bash"
    user "root"

	search("aws_opsworks_layer").each do |layer|
		if layer['name'] == 'queue'
			instances = search("aws_opsworks_instance", "layer_ids:" + "#{layer['layer_id']}"  + ' AND status:online')
			instances.each do |instance|
				Chef::Log.warn("********** The instance's ip is '#{instance['private_ip']}' **********")
				code <<-EOH
      				docker run -d --name=#{name} #{image} -e RABBITMQ_SERVER='#{instance['private_ip']}' 
    			EOH
			end
		end
	end
  end
end
