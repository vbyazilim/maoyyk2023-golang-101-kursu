task :command_exists, [:command] do |_, args|
  abort "#{args.command} doesn't exists" if `command -v #{args.command} > /dev/null 2>&1 && echo $?`.chomp.empty?
end

task :has_bumpversion do
  Rake::Task['command_exists'].invoke('bumpversion')
end

task :is_repo_clean do
  abort 'please commit your changes first!' unless `git status -s | wc -l`.strip.to_i.zero?
end

AVAILABLE_REVISIONS = %w[major minor patch].freeze
task :bump, [:revision] => [:has_bumpversion] do |_, args|
  args.with_defaults(revision: 'patch')
  unless AVAILABLE_REVISIONS.include?(args.revision)
    abort "Please provide valid revision: #{AVAILABLE_REVISIONS.join(',')}"
  end

  system "bumpversion #{args.revision}"
  exit $?.exitstatus
end

desc "release new version #{AVAILABLE_REVISIONS.join(',')}, default: patch"
task :release, [:revision] => [:is_repo_clean] do |_, args|
  args.with_defaults(revision: 'patch')
  Rake::Task['bump'].invoke(args.revision)
end

namespace :mkdocs do
  desc "run docs server"
  task :serve do
    system "mkdocs serve"
  end

  desc "build docs"
  task :build do
    system "mkdocs build --clean"
  end

  desc "deploy to GitHub"
  task :deploy do
    system "mkdocs gh-deploy"
  end
end
