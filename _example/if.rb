directory '/tmp/a' do
  not_if 'true'
end

execute 'true' do
  not_if 'false'
  user 'root'
  cwd '/tmp'
end

file '/tmp/b' do
  only_if 'false'
end
