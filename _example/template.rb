# itamae-go local -j _example/files/node.json _example/template.rb

file '/tmp/hello' do
  content 'hello world'
  action :create
end

template '/tmp/template_result' do
  source '_example/files/template.erb'
  action :create
end
