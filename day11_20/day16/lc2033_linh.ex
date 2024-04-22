defmodule Solution do
  @spec min_operations(grid :: [[integer]], x :: integer) :: integer
  def min_operations(grid, x) do
    l = grid |> List.flatten()
    r = rem(List.first(l), x)
    Enum.all?(l, fn i -> rem(i, x) == r end)
    |> case do
      false -> -1
      true  ->
        m = median(l)
        l |> Enum.map(fn i -> abs(div(i-r, x) - div(m, x)) end) |> Enum.sum()
    end
  end

  defp median(list) do
    i = list |> length() |> div(2)
    list |> Enum.sort() |> Enum.at(i)
  end
end
