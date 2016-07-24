include_recipe 'tmux.rb'

define :vim, options: '--with-lua --with-luajit' do
  ver  = params[:name]
  opts = params[:options]
  package 'vim' do
    version ver
    options opts
  end
end

vim '7.4.1910-1'
