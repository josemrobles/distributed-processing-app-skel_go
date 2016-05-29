search("aws_opsworks_layer").each do |layer|
	if layer['name'] == 'queue'
		instances = search("aws_opsworks_instance", "layer_ids:" + "#{layer['layer_id']}"  + ' AND status:online')
		instances.each do |instance|
			Chef::Log.warn("********** The instance's ip is '#{instance['private_ip']}' **********")
		end
	end
end
