const std=@import("std");
const Allocator = std.mem.Allocator;

pub fn main() void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();

    const input = readInput(allocator);
    std.debug.print("{s}\n", .{input});
}

fn readInput(allocator: Allocator) []const u8 {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();
    const contents = try file.readToEndAlloc(allocator, 16384);
    return contents;
}