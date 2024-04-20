defmodule Solution do
  @spec min_operations(grid :: [[integer]], x :: integer) :: integer
  def min_operations(grid, x) do
    l = grid |> List.flatten()
    r = rem(List.first(l), x)
    Enum.all?(l, fn i -> rem(i, x) == r end)
    |> case do
      false -> -1
      true  ->
        l = Enum.map(l, fn i -> div(i-r, x) end)
        m = median(l)
        Enum.map(l, fn i -> abs(i - m) end)
        |> Enum.reduce(fn i, j -> i + j end)
    end
  end

  defp median(list) do
    i = list |> length() |> div(2)
    list |> Enum.sort() |> Enum.at(i)
  end
end
