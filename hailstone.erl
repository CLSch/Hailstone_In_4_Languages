% Caitlin Schäffers
% Programmeertalen
% University of Amsterdam
%
% Make a hailstone function with coroutines.

-module(hailstone).

-export([main/0]).

main() ->
    % get start value from user
    StartValue =
	list_to_integer(io:get_line("Give your number: ") --
			  "\n"),
    Seq = hailstone_producer([StartValue]),
    % send all the calculated values to the consumer
    lists:foreach(fun (N) -> self() ! N end,
		  lists:reverse(Seq)),
    io:fwrite("Done producing!\n"),
    hailstone_consumer().

% calculate next value in sequence and add to list
hailstone_producer([1 | T]) -> [1 | T];
hailstone_producer([N | T]) when N rem 2 == 0 ->
    hailstone_producer([N div 2, N | T]);
hailstone_producer([N | T]) ->
    hailstone_producer([N * 3 + 1, N | T]).

% a potential other way to calculate the sequence:

% hailstone_producer(1) -> [1];
% hailstone_producer(N) when N rem 2 == 0 ->
% 	[N|hailstone_producer(N div 2)];
% hailstone_producer(N) ->
% 	[N|hailstone_producer(N * 3 + 1)].

hailstone_consumer() ->
    io:get_line("Press Enter to receive the next value."),
    % receive whatever is being send in the program
    receive
      1 ->
	  io:format("Received: 1, the end of the sequence.~n");
      N ->
	  io:format("Received: ~b~n", [N]), hailstone_consumer()
    end.
