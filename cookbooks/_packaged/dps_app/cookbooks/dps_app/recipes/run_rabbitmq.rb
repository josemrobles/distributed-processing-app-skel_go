node[:apps].each do |name, image|  
  script "run_app_#{name}_container" do
    interpreter "bash"
    user "root"
    code <<-EOH
      docker run -d -p 9002:15672 -p 5672:5672 --name=#{name} #{image}
    EOH
  end
end
