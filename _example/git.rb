git '/tmp/Spoon-Knife' do
  repository 'https://github.com/octocat/Spoon-Knife.git'
  action :sync
end

directory '/tmp/Spoon-Knife' do
  action :delete
end
