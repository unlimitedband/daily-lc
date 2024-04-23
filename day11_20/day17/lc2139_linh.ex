defmodule Solution do
  @spec min_moves(target :: integer, max_doubles :: integer) :: integer
  def min_moves(target, max_doubles) do
    move(target, max_doubles, 0)
  end

  defp move(target, 0, acc), do: acc + target - 1
  defp move(1, _, acc), do: acc
  defp move(target, max_doubles, acc) do
    case rem(target, 2) do
      0 -> move(div(target, 2), max_doubles - 1, acc + 1)
      1 -> move(target - 1, max_doubles, acc + 1)
    end
  end
end
