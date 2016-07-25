file '/tmp/hello' do
  content 'hello world'
  action :create
end

file '/tmp/world' do
  content 'world is over'
end

link '/tmp/world' do
  to '/tmp/hello'
  force 'true'
end
