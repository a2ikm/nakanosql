#!/usr/bin/env ruby

def assert_equal(expected, actual)
  unless expected == actual
    backtrace = caller(1, 1).first
    abort "#{backtrace}: expected \"#{expected}\" but got \"#{actual}\""
  end
end

system "rm data.db"
system "echo \"INSERT 1\" | ./nakanosql"
system "echo \"INSERT 2 3\" | ./nakanosql"
system "echo \"INSERT 4 5 6\" | ./nakanosql"
result = `echo \"SELECT\" | ./nakanosql`
assert_equal "1\n2\t3\n4\t5\t6\n", result
