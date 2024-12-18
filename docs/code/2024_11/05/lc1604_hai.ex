defmodule Solution do
  @spec alert_names(key_name :: [String.t], key_time :: [String.t]) :: [String.t]
  def alert_names(key_name, key_time) do
    arr = Enum.zip(key_name, Enum.map(key_time, fn x -> time_to_minutes(x) end))
    arr = Enum.sort(arr, fn {a1, b1}, {a2, b2} ->
        if a1 == a2 do
            b1 <= b2
        else
            a1 <= a2
        end
    end)

    maped = Enum.reduce(arr, %{}, fn {k, v}, acc ->
        Map.update(acc, k, [v], fn cur ->
            [v | cur]
        end)
    end)

    Map.keys(maped)
    |> Enum.reduce([], fn key, acc ->
        arr = maped[key]
        mminus = if length(arr) > 2 do
            Enum.map(Enum.zip(Enum.slice(arr, 0, length(arr) - 2), Enum.slice(arr, 2, length(arr) - 2)), fn {x, y} -> x - y end) |> Enum.min()
            else
                61
            end
        if  mminus <= 60 do
            acc ++ [key]
        else
            acc
        end
    end) |> Enum.sort()


  end

  def time_to_minutes(time_string) do
    [hours, minutes] = String.split(time_string, ":") |> Enum.map(&String.to_integer/1)
    hours * 60 + minutes
  end
end
