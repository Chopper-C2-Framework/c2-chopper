import { create } from "zustand";

interface SearchState {
    search: string;
    updateSearchField: (newValue: string) => void;
    clearSearchField: () => void;
}

const usePluginsSearchStore = create<SearchState>((set) => ({
  search: "",
  updateSearchField: (newValue: string) =>
    set((_) => ({ search: newValue })),
  clearSearchField: () => set((_) => ({ search: "" })),
}));


export default usePluginsSearchStore;