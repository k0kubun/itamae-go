service 'sshd' do
  action [:start, :stop, :enable, :disable]
end
