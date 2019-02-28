#!/usr/bin/env ruby

def assert_equal(expected, actual)
  unless expected == actual
    backtrace = caller(1, 1).first
    abort "#{backtrace}: expected \"#{expected}\" but got \"#{actual}\""
  end
end

def test(command, expected)
  assert_equal expected, %x{#{command}}.chomp
end

test "echo \"SELECT 1\" | ./nakanosql", "SELECT 1"
